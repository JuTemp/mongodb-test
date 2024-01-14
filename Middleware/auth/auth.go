package auth

import (
	"net/http"

	"github.com/JuTemp/mongodb-test/Logical/auth"
	"github.com/JuTemp/mongodb-test/Middleware/encrypt"
	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

var Check gin.HandlerFunc = func(ctx *gin.Context) {
	token := ctx.GetHeader("X-Auth-Token")
	if !auth.CheckToken(token) {
		MyGinContext.MyGinContext{Context: ctx}.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
			Code:            http.StatusForbidden,
			UnencryptedJSON: gin.H{"message": "Token is invalid"},
		})
		encrypt.Encrypt(ctx)
		ctx.Abort()
		return
	}

	ctx.Next()
}
