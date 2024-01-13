package auth

import (
	"github.com/JuTemp/mongodb-test/util/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckToken(token string) bool {
	// Don't spread out the token in mongodb
	tokenCorrect, result := mongodb.FindOne[Token](Collection, bson.D{{Key: "token", Value: token}}, bson.D{{Key: "token", Value: 1}})
	if !result {
		return false
	}
	if tokenCorrect.Token == token {
		return true
	}
	return false
}
