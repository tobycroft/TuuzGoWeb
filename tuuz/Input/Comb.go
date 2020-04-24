package Input

import (
	"github.com/gin-gonic/gin"
)

func Combi(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := Post(key, c, xss)
	if !ok {
		return Get(key, c, xss)
	} else {
		return in, true
	}
}
