package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find[T any](client *mongo.Collection, filter bson.D, projection bson.D) ([]T, error) {
	cur, err := client.Find(Context, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	var result []T
	cur.All(Context, &result)
	return result, nil
}
