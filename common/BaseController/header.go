package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"main.go/tuuz/RET"
)

func header_handler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("S-P-I", c.ClientIP())
	c.Header("S-P-H", c.Request.Host)
	c.Header("S-P-P", app_conf.Project)
	c.Header("S-P-M", app_conf.AppMode)
	c.Header("S-V-L", app_conf.Version_lowest)
	c.Header("S-V-C", app_conf.Version_current)
	c.Header("S-V-ED", app_conf.Version_end_date)
}

func header_auth(c *gin.Context) (ok bool, uid string, token string, debug string) {
	uid = c.GetHeader("uid")
	if uid == "" {
		c.JSON(RET.Ret_fail(-1, nil, "Header-[uid]"))
		return
	}
	token = c.GetHeader("token")
	if token == "" {
		c.JSON(RET.Ret_fail(-1, nil, "Header-[token]"))
		return
	}
	debug = c.GetHeader("debug")
	ok = true
	return
}
