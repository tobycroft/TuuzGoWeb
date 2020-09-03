package Input

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"main.go/tuuz/RET"
)

func Get(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetQuery(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "GET-["+key+"]"))
		c.Abort()
		return "", false
	} else {
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}

func GetBool(key string, c *gin.Context) (bool, bool) {
	in, ok := c.GetQuery(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "GET-["+key+"]"))
		c.Abort()
		return false, false
	} else {
		switch in {
		case "1":
			return true, true

		case "0":
			return false, true

		case "true":
			return true, true

		case "false":
			return false, true

		default:
			return false, false
		}
	}
}
