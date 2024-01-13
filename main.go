package main

import (
	"github.com/JuTemp/mongodb-test/Api"
	"github.com/JuTemp/mongodb-test/Middleware/auth"
	"github.com/JuTemp/mongodb-test/Middleware/encrypt"
	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine = gin.Default()
	r.GET("/ping", func(ctx *gin.Context) { Api.Ping(ctx) })
	r.POST("/getToken", func(ctx *gin.Context) { Api.GetToken(ctx) })

	authGroup := r.Group("/auth")
	authGroup.Use(auth.Check, encrypt.Encrypt)
	authGroup.POST("/search", func(ctx *gin.Context) { Api.Search(ctx) })
	r.Run(":38576") // listen and serve on 0.0.0.0:38576
}
