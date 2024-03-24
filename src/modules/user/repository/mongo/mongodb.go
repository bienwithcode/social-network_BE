package rMongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type mongodbStorage struct {
	db *mongo.Database
}

func NewMongoStorage(db *mongo.Database) *mongodbStorage {
	return &mongodbStorage{db: db}
}
