package encrypt

import (
	MyGinContext "github.com/JuTemp/mongodb-test/util/myGinContext"
	"github.com/gin-gonic/gin"
)

var Encrypt = func(ctx *gin.Context) {
	ctx.Next()

	cau := MyGinContext.MyGinContext{Context: ctx}.GetCodeAndUnencryptedJSON()

	ctx.JSON(cau.Code, gin.H{"inner": cau.UnencryptedJSON})
}
