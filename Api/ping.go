package Api

import (
	"net/http"

	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

func Ping(ctx MyGinContext.MyGinContext) {
	ctx.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
		Code:            http.StatusOK,
		UnencryptedJSON: gin.H{"message": "pong"},
	})
}
