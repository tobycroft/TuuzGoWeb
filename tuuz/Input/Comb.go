package Input

import (
	"github.com/gofiber/fiber/v2"
	"main.go/tuuz/RET"
)

func Combi(key string, c *fiber.Ctx, xss bool) (string, bool) {
	in := c.FormValue(key)
	if len(in) == 0 {
		in = c.Query(key)
		if len(in) == 0 {
			c.JSON(RET.Ret_fail(400, key, key))
			return "", false
		} else {
			return in, true
		}
	} else {
		return in, true
	}
}
