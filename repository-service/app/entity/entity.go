package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Repo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RepoId      int64              `json:"repo_id" bson:"repo_id"`
	Name        string             `json:"name" bson:"name"`
	Owner       string             `json:"owner" bson:"owner"`
	OwnerId     int64              `json:"owner_id" bson:"owner_id"`
	Description string             `json:"description" bson:"description"`
	Language    string             `json:"language" bson:"language"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type SearchResult struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Query     string             `json:"query" bson:"query"`
	Results   []*Repo            `json:"query" bson:"results"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
}

func (s *SearchResult) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
