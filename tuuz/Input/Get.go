package Input

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/RET"
)

func Get(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetQuery(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, key))
		c.Abort()
		return "", false
	} else {
		return in, true
	}
}
