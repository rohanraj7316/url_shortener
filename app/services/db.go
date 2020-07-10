package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type collectionList struct {
	Analytics *mongo.Collection
	URLData   *mongo.Collection
}

// DbConnection mongo connection
var DbConnection *mongo.Client
var c collectionList

// MongoConnection establish connection to mongo db
func MongoConnection(config interface{}) error {
	// TODO: get this value from env.
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	DbConnection = client
	return nil
}

// MongoConnectionHealthCheck connection ping
func MongoConnectionHealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := DbConnection.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func (c collectionList) initializeCollection(name string) *collectionList {
	db := DbConnection.Database(name)
	c.Analytics = db.Collection("Analytics")
	c.URLData = db.Collection("URLData")

	return &c
}

// GetCollection get collection
func GetCollection(collection string) *mongo.Collection {
}
