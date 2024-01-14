package Api

import (
	"net/http"

	"github.com/JuTemp/mongodb-test/Logical/auth"
	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

func GetToken(c MyGinContext.MyGinContext) {
	username := c.PostForm("username")
	token, result := auth.NewToken(username)
	if !result {
		c.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
			Code:            http.StatusForbidden,
			UnencryptedJSON: gin.H{"error": "Cannot new token"},
		})
		return
	}
	c.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
		Code:            http.StatusOK,
		UnencryptedJSON: gin.H{"token": token},
	})
}
