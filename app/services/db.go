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

// DBConnection mongo connection
var DBConnection *mongo.Client

// DB mongo database
var DB *mongo.Database

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
	DBConnection = client
	return nil
}

// MongoConnectionHealthCheck connection ping
func MongoConnectionHealthCheck() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := DBConnection.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	// TODO: need to pull DB name from env
	DB = DBConnection.Database("DemoDB")

	return nil
}

// GetCollection get collection
func getCollection(c string) *mongo.Collection {
	if c == "Analytics" {
		return DB.Collection(c)
	} else if c == "URLData" {
		return DB.Collection(c)
	} else {
		return nil
	}
}

// UpdateData update data into given collection
func UpdateData(ctx context.Context, c string, filter interface{}, update interface{}) (*mongo.SingleResult, error) {
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

	result := collection.FindOne(ctx, filter, opt)
	if err := result.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// InsertData insert data into given collection
func InsertData(ctx context.Context, c string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := getCollection(c)

	var opt *options.InsertOneOptions
	opt.SetBypassDocumentValidation(false)

	result, err := collection.InsertOne(ctx, doc, opt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
