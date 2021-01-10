package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginedController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("S-P-I", c.ClientIP())
		c.Header("S-P-P", app_conf.Project)
		c.Header("S-P-M", app_conf.AppMode)
		uid, ok := Input.Post("uid", c, false)
		if !ok {
			c.Abort()
			return
		}
		token, ok := Input.Post("token", c, false)
		if !ok {
			c.Abort()
			return
		}
		debug, ok := c.GetPostForm("debug")
		if ok {
			if debug == app_conf.Debug && app_conf.TestMode {
				c.Next()
				return
			}
		}
		if len(TokenModel.Api_find(uid, token)) > 0 {
			c.Next()
			return
		} else {
			RET.Fail(c, -1, "Auth_fail", "未登录")
			c.Abort()
			return
		}
	}
}

func LoginWSController() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("S-P-I", c.ClientIP())
		c.Header("S-P-P", app_conf.Project)
		c.Header("S-P-M", app_conf.AppMode)
		uid, ok := Input.Post("uid", c, false)
		if !ok {
			c.Abort()
			return
		}
		ws, ok := c.GetPostForm("wskey")
		if ok {
			if ws == app_conf.WebsocketKey {
				c.Next()
				return
			}
		}
		token, ok := Input.Post("token", c, false)
		if !ok {
			c.Abort()
			return
		}
		debug, ok := c.GetPostForm("debug")
		if ok {
			if debug == app_conf.Debug && app_conf.TestMode {
				c.Next()
				return
			}
		}
		if len(TokenModel.Api_find(uid, token)) > 0 {
			c.Next()
			return
		} else {
			RET.Fail(c, -1, "Auth_fail", "未登录")
			c.Abort()
			return
		}
	}
}
