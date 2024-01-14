package main

import (
	"github.com/JuTemp/mongodb-test/Api"
	"github.com/JuTemp/mongodb-test/Middleware/auth"
	"github.com/JuTemp/mongodb-test/Middleware/encrypt"
	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine = gin.Default()
	r.Use(encrypt.Encrypt)
	r.GET("/ping", func(ctx *gin.Context) { Api.Ping(MyGinContext.MyGinContext{Context: ctx}) })
	r.POST("/getToken", func(ctx *gin.Context) { Api.GetToken(MyGinContext.MyGinContext{Context: ctx}) })

	authGroup := r.Group("/auth")
	authGroup.Use(auth.Check, encrypt.Encrypt) // Double encrypt
	authGroup.POST("/search", func(ctx *gin.Context) { Api.Search(MyGinContext.MyGinContext{Context: ctx}) })
	r.Run(":38576") // listen and serve on 0.0.0.0:38576
}
