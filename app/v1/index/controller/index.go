package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("/", index)
	route.Any("/login", loginss)
	route.Any("/register")
}

func index(c *gin.Context) {
	c.String(0, "index")
}

func loginss(c *gin.Context) {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(0, json)
}
