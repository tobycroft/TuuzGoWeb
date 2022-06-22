package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/RET"
	"net/http"
)

func LoginedController() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		header_handler(c)
		uid := ""
		token := ""
		debug := ""
		ok := false
		if app_conf.HeaderAuthMode {
			ok, uid, token, debug = header_auth(c)
			if !ok {
				c.Abort()
				return
			}
		} else {
			ok, uid, token, debug = post_auth(c)
			if !ok {
				c.Abort()
				return
			}
		}
		if app_conf.TestMode {
			if debug == app_conf.Debug {
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

func post_auth(c *gin.Context) (ok bool, uid string, token string, debug string) {
	uid = c.GetHeader("uid")
	if len(uid) < 1 {
		c.JSON(RET.Ret_fail(-1, nil, "Header-[uid]"))
		return false, "", "", ""
	}
	token = c.GetHeader("token")
	if len(token) < 1 {
		c.JSON(RET.Ret_fail(-1, nil, "Header-[token]"))
		return false, "", "", ""
	}
	debug = c.GetHeader("debug")
	return
}

func LoginWSController() gin.HandlerFunc {
	return func(c *gin.Context) {
		header_handler(c)
		uid := c.GetHeader("uid")
		if len(uid) < 1 {
			c.Abort()
			return
		}
		ws := c.GetHeader("wskey")
		if len(uid) < 1 {
			if ws == app_conf.WebsocketKey {
				c.Next()
				return
			}
		}
		token := c.GetHeader("token")
		if len(token) < 1 {
			c.Abort()
			return
		}
		debug := c.GetHeader("debug")
		if len(token) < 1 {
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
