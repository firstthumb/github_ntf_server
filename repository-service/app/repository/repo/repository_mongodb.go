package repo

import (
	"context"
	"errors"
	"githubntf/app/config"
	"githubntf/app/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const RepositoryCollection = "repositories"
const SearchResultCollection = "search_results"
const CacheExpireInHour = 1

type MongoRepository struct {
	pool *mongo.Database
	db   string
}

func NewMongoRepository(c *config.Config) *MongoRepository {
	mongoUrl, err := config.GetSecret(c.Secret.DB)
	if err != nil {
		log.Fatal("Couldn't fetch mongo secret: %v", err)
		return nil
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))

	if err != nil {
		log.Fatal("Couldn't connect to mongo: %v", err)
		return nil
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal("Mongo client couldn't connect with background context: %v", err)
		return nil
	}

	log.Println("Connected to database...")

	return &MongoRepository{
		pool: client.Database("test"),
		db:   "test",
	}
}

func (r *MongoRepository) Find(id string) (*entity.Repo, error) {
	result := entity.Repo{}
	coll := r.pool.Collection(RepositoryCollection)
	err := coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	switch err {
	case nil:
		return &result, nil
	default:
		return nil, err
	}
}

func (r *MongoRepository) Store(b *entity.Repo) (string, error) {
	coll := r.pool.Collection(RepositoryCollection)
	res, err := coll.InsertOne(context.TODO(), b)
	if err != nil {
		log.Printf("Failed to insert: %s", err)
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *MongoRepository) FindSearchResult(query string) ([]*entity.Repo, error) {
	var searchResult *entity.SearchResult
	coll := r.pool.Collection(SearchResultCollection)
	err := coll.FindOne(context.TODO(), bson.M{"query": query}).Decode(&searchResult)
	if err != nil {
		log.Printf("Could not find any search result : %s", err)
		return nil, err
	}

	if searchResult.IsExpired() {
		log.Printf("Search result expired")
		return nil, errors.New("search result expired")
	}

	return searchResult.Results, nil
}

func (r *MongoRepository) StoreSearchResult(query string, results []*entity.Repo) error {
	searchResult := entity.SearchResult{
		Query:     query,
		Results:   results,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(CacheExpireInHour)),
	}

	coll := r.pool.Collection(SearchResultCollection)
	updateResult, _ := coll.ReplaceOne(context.TODO(), bson.M{"query": query}, searchResult)
	if updateResult.ModifiedCount == 0 {
		_, err := coll.InsertOne(context.TODO(), searchResult)

		if err != nil {
			log.Printf("Failed to store search result: %s", err)
			return errors.New("failed to store search result")
		}
	}

	return nil
}
