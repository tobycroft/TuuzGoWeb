package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz/Input"
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
	file, ok := Input.Upload(c)
	if !ok {
		return
	}
	fmt.Println(file)
}

type TokenRequest struct {
	Project string `json:"project" binding:"required"`
	Secret  string `json:"secret" binding:"required"`
}

func login_json(c *gin.Context) {
	var tk TokenRequest
	c.ShouldBindWith(tk, Input.JsonHS)
	json := map[string]string{}
	json["username"] = "123"
	json["password"] = "123"
	c.JSON(0, json)
}
