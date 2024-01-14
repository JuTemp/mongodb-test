package MyGinContext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyGinContext struct {
	*gin.Context
}

type CodeAndUnencryptedJSON struct {
	Code            int   `json:"Code"`
	UnencryptedJSON gin.H `json:"UnencryptedJSON"`
}

const keyOfMyGinContext = "CodeAndUnencryptedJSON"

func (ctx MyGinContext) SetCodeAndUnencryptedJSON(cau CodeAndUnencryptedJSON) {
	ctx.Set(keyOfMyGinContext, cau)
}

func (ctx MyGinContext) GetCodeAndUnencryptedJSON() CodeAndUnencryptedJSON {
	value, exist := ctx.Get(keyOfMyGinContext)
	if !exist {
		return CodeAndUnencryptedJSON{
			Code:            http.StatusBadRequest,
			UnencryptedJSON: gin.H{},
		}
	}
	return value.(CodeAndUnencryptedJSON)
}
