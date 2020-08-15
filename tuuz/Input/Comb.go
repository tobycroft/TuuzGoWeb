package Input

import (
	"github.com/gin-gonic/gin"
	"main.go/tuuz/RET"
)

func Combi(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		in, ok = c.GetQuery(key)
		if !ok {
			c.JSON(RET.Ret_fail(400, key, key))
			c.Abort()
			return "", ok
		} else {
			return in, ok
		}
	} else {
		return in, ok
	}
}
