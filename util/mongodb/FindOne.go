package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](client *mongo.Collection, filter interface{}, projection interface{}) (T, bool) {
	var result T
	var err error
	if projection == nil {
		err = client.FindOne(Context, filter).Decode(&result)
	} else {
		err = client.FindOne(Context, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	}
	return result, err == nil
}
