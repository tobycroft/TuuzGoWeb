package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/common/BaseModel/TokenModel"
	"main.go/config/app_conf"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
	"net/http"
)

func LoginedController() gin.HandlerFunc {
	return func(c *gin.Context) {
		header_handler(c)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
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
	uid, ok = c.GetPostForm("uid")
	if !ok {
		c.JSON(RET.Ret_fail(-1, nil, "POST-[uid]"))
		return
	}
	token, ok = c.GetPostForm("token")
	if !ok {
		c.JSON(RET.Ret_fail(-1, nil, "POST-[token]"))
		return
	}
	debug = c.PostForm("debug")
	return
}

func LoginWSController() gin.HandlerFunc {
	return func(c *gin.Context) {
		header_handler(c)
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
