package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz/Input"
	"time"
)

func IndexController(route *gin.RouterGroup) {
	route.Any("", index)
	route.Any("login", loginss)
	route.Any("upload", upload)
	route.Any("register")
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

func upload(c *gin.Context) {
	fmt.Println(time.Now().Format("200603"))
	file, ok := Input.PostFile(c)
	if !ok {
		return
	}
	fmt.Println(file)
}
