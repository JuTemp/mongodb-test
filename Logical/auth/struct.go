package auth

type Token struct {
	Username string `bson:"username"`
	Token    string `bson:"token"`
}
