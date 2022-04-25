package RET

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"main.go/config/app_conf"
	"main.go/tuuz/Jsong"
)

func Json(data interface{}) string {
	ret, _ := Jsong.Encode(data)
	return ret
}

func Success(c *fiber.Ctx, code int, data, echo interface{}) {
	if echo == nil {
		switch code {

		case 0:
			echo = "成功"
			break

		case -1:
			echo = "登录失效请重新登录"
			break

		case 400:
			echo = "参数错误"
			break

		case 401:
			echo = "鉴权失败"
			break

		case 403:
			echo = "权限不足"
			break

		case 406, 407:
			echo = "数据不符合期待"
			break

		case 404:
			echo = "未找到数据"
			break

		case 500:
			echo = "数据库错误"
			break

		default:
			echo = "失败"
			break
		}
	}
	switch echo.(type) {
	case error:
		echo = echo.(error).Error()
		break
	default:
		break
	}
	if data == nil {
		data = []interface{}{}
	}
	if app_conf.SecureJson {
		c.JSON(Ret_succ(code, data, echo))
	} else {
		c.JSON(Ret_succ(code, data, echo))
	}
	return
}

func Fail(c *fiber.Ctx, code int, data, echo interface{}) {
	Success(c, code, data, echo)
	return
}

func Ret_succ(code int, data, echo interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	if data == nil {
		data = []interface{}{}
	}
	ret["code"] = code
	ret["data"] = data
	ret["echo"] = echo
	return ret
}

func Ret_fail(code int, data, echo interface{}) map[string]interface{} {
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
