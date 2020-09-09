package Input

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"html/template"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
)

func Post(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return "", false
	} else {
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}

func PostInt(key string, c *gin.Context) (int, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Int(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, key+" should be int", key+" should be int"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func PostInt64(key string, c *gin.Context) (int64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Int64(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, key+" should be int64", key+" should be int64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func PostFloat64(key string, c *gin.Context) (float64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return 0, false
	} else {
		i, e := Calc.String2Float64(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, key+" should be float64", key+" should be float64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func PostBool(key string, c *gin.Context) (bool, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return false, false
	} else {
		switch in {
		case "1":
			return true, true

		case "0":
			return false, true

		case "true":
			return true, true

		case "false":
			return false, true

		default:
			return false, false
		}
	}
}

func PostArray(key string, c *gin.Context) ([]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return nil, false
	} else {
		i, err := Jsong.JArray(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-Array", key+" should be a Json-Array"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func PostObject(key string, c *gin.Context) (map[string]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return nil, false
	} else {
		i, err := Jsong.JObject(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-Object", key+" should be a Json-Object"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func PostArrayObject(key string, c *gin.Context) ([]map[string]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return nil, false
	} else {
		i, err := Jsong.JArrayObject(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-ArrayObject", key+" should be a Json-ArrayObject"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func PostAny(key string, c *gin.Context, AnyType interface{}) bool {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return false
	} else {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err := json.Unmarshal([]byte(in), &AnyType)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-AnyType", key+" should be a Json-AnyType"))
			c.Abort()
			return false
		}
		return true
	}
}
