package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Repo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
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
