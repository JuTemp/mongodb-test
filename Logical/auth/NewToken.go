package auth

import (
	"github.com/JuTemp/mongodb-test/util/mongodb"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func NewToken(username string) (string, bool) {
	token := uuid.New().String()
	_, result := mongodb.FindOne[Token](Collection, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "username", Value: username}})
	if result {
		result = mongodb.ReplaceOne(Collection, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "username", Value: username}, {Key: "token", Value: token}})
	} else {
		result = mongodb.InsertOne(Collection, bson.D{{Key: "username", Value: username}, {Key: "token", Value: token}})
	}
	if !result {
		return "", false
	}
	return token, true
}
