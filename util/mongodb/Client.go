package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = getClient("mongodb://localhost:27017", Context)

func getClient(uri string, ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func GetDatabase(dbName string) *mongo.Database {
	return Client.Database(dbName)
}

func GetCollection(dbName string, collName string) *mongo.Collection {
	return GetDatabase(dbName).Collection(collName)
}
