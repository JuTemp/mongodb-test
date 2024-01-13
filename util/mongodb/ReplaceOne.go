package mongodb

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func ReplaceOne(coll *mongo.Collection, filter interface{}, replacement interface{}) bool {
	_, err := coll.ReplaceOne(Context, filter, replacement)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
