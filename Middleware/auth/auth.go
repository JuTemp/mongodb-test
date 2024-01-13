package auth

import (
	"log"
	"net/http"

	"github.com/JuTemp/mongodb-test/Logical/auth"
	"github.com/gin-gonic/gin"
)

var Check gin.HandlerFunc = func(c *gin.Context) {
	token := c.GetHeader("X-Auth-Token")
	if !auth.CheckToken(token) {
		log.Println("failed")
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Token is invalid"})
	}
}
