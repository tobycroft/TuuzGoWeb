package RET

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/tuuz/Jsong"
)

func Json(data interface{}) string {
	ret, _ := Jsong.Encode(data)
	return ret
}

func Success(c *gin.Context, code int, data, echo interface{}) {
	c.JSON(Ret_succ(code, data, echo))
	c.Abort()
	return
}

func Fail(c *gin.Context, code int, data, echo interface{}) {
	Success(c, code, data, echo)
	return
}

func Ret_succ(code int, data, echo interface{}) (int, map[string]interface{}) {
	ret := make(map[string]interface{})
	ret_code := -1
	if code == 0 {
		ret_code = 200
	} else {
		ret_code = code
	}
	ret["code"] = code
	ret["data"] = data
	ret["echo"] = echo
	return ret_code, ret
}

func Ret_fail(code int, data, echo interface{}) (int, map[string]interface{}) {
	return Ret_succ(code, data, echo)
}

func Ws_succ(typ string, code interface{}, data interface{}, echo interface{}) string {
	ret := make(map[string]interface{})
	ret["type"] = typ
	ret["code"] = code
	ret["data"] = data
	ret["echo"] = echo
	jb, err := Jsong.Encode(ret)
	//fmt.Println(jb)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jb)
}

func Ws_succ2(typ string, route string, code interface{}, data interface{}, echo interface{}) string {
	ret := make(map[string]interface{})
	ret["type"] = typ
	ret["route"] = route
	ret["code"] = code
	ret["data"] = data
	ret["echo"] = echo
	jb, err := Jsong.Encode(ret)
	//fmt.Println(jb)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jb)
}

func Ws_fail(typ string, code interface{}, data interface{}, echo interface{}) string {
	return Ws_succ(typ, code, data, echo)
}
