package ginreturn

import (
	"github.com/gin-gonic/gin"
)

func GinReturn(c *gin.Context, code int, data interface{}, err error) {
	c.JSON(code, genResponse(data, err))
}

func genResponse(data interface{}, err error) map[string]interface{} {
	res := make(map[string]interface{})
	res["data"] = data
	if err != nil {
		res["msg"] = err.Error()
	} else {
		res["msg"] = ""
	}

	return res
}
