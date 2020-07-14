package services

import (
	"context"
	"math/rand"
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
func MongoConnection() error {
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
func getCollection(collection string) *mongo.Collection {
	// TODO: need to figure out way to return the object only
	return nil
}

// InsertData insert data into given collection
func InsertData(ctx context.Context, c string, filter interface{}, update interface{}) (*mongo.SingleResult, error) {
	collection := getCollection(c)

	var opt options.FindOneAndUpdateOptions
	opt.SetUpsert(true)
	opt.SetMaxTime(time.Duration(rand.Int31n(5000)) * time.Millisecond)

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if err := result.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// FindData find and return data from collection
func FindData(ctx context.Context, c string, filter interface{}) (*mongo.SingleResult, error) {
	collection := getCollection(c)

	var opt *options.FindOneOptions

	res := collection.FindOne(ctx, filter, opt)
	if err := res.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
