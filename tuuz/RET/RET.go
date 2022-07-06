package RET

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"main.go/tuuz/Jsong"
	"net/http"
)

func Json(data interface{}) string {
	ret, _ := Jsong.Encode(data)
	return ret
}

func Success(c *gin.Context, code int, data, echo interface{}) {
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
		retCode, retJson := Ret_succ(code, data, echo)
		secure_json(c, &retCode, &retJson)
	} else {
		retCode, retJson := Ret_succ(code, data, echo)
		json(c, &retCode, &retJson)
	}
	c.Abort()
	return
}

func writeContentType(w http.ResponseWriter,
	value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
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
		ret_code = 200
	}
	if data == nil {
		data = []interface{}{}
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
