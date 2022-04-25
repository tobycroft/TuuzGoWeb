package Input

import (
	"github.com/gofiber/fiber/v2"
	"html/template"
	"main.go/tuuz/RET"
)

func Get(key string, c *fiber.Ctx, xss bool) (string, bool) {
	in := c.Query(key)
	if len(in) == 0 {
		c.JSON(RET.Ret_fail(400, key, "GET-["+key+"]"))
		return "", false
	} else {
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}

func GetBool(key string, c *fiber.Ctx) (bool, bool) {
	in := c.Query(key)
	if len(in) == 0 {
		c.JSON(RET.Ret_fail(400, key, "GET-["+key+"]"))
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
