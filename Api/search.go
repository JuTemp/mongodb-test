package Api

import (
	"net/http"

	Logical "github.com/JuTemp/mongodb-test/Logical"
	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

func Search(c MyGinContext.MyGinContext) {
	username := c.PostForm("username")
	if username == "" {
		c.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
			Code:            http.StatusBadRequest,
			UnencryptedJSON: gin.H{"error": "Username is required"},
		})
		return
	}
	note, result := Logical.Search(username)
	if !result {
		c.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
			Code:            http.StatusInternalServerError,
			UnencryptedJSON: gin.H{"error": "no search"},
		})
		return
	}
	c.SetCodeAndUnencryptedJSON(MyGinContext.CodeAndUnencryptedJSON{
		Code:            http.StatusOK,
		UnencryptedJSON: gin.H{"note": note},
	})
}
