package BaseController

import (
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
)

func header_handler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("S-P-I", c.ClientIP())
	c.Header("S-P-P", app_conf.Project)
	c.Header("S-P-M", app_conf.AppMode)
	c.Header("S-V-L", app_conf.Version_lowest)
	c.Header("S-V-C", app_conf.Version_current)
	c.Header("S-V-ED", app_conf.Version_end_date)
}
