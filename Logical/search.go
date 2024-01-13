package logical

import (
	"github.com/JuTemp/mongodb-test/util/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func Search(username string) (string, bool) {
	user, result := mongodb.FindOne[User](Collection,
		bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "note", Value: 1}})
	if !result {
		return "", false
	}
	return user.Note, true
}

type User struct {
	ID       int    `bson:"id"`
	Username string `bson:"username"`
	Note     string `bson:"note"`
}
