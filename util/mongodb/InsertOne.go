package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(coll *mongo.Collection, doc interface{}) bool {
	_, err := coll.InsertOne(Context, doc)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
