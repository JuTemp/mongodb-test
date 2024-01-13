package Api

import (
	"net/http"

	"github.com/JuTemp/mongodb-test/Logical/auth"
	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	username := c.PostForm("username")
	token, result := auth.NewToken(username)
	if !result {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot new token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
