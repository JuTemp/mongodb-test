package Api

import (
	"net/http"

	logical "github.com/JuTemp/mongodb-test/Logical"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}
	note, result := logical.Search(username)
	if !result {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no search"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"note": note})
}
