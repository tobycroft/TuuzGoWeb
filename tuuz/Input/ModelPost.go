package Input

import (
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/tobycroft/gorose-pro"
	"html/template"
	"main.go/tuuz/RET"
	"main.go/tuuz/Vali"
	"time"
)

func MPostAuto(data *gorose.Data, c *gin.Context, xss bool) {
	for key, _ := range *data {
		MPost(key, data, c)
	}
	return
}

func MPost(key string, data *gorose.Data, c *gin.Context) (ok bool) {
	in, ok := c.GetPostForm(key)
	temp_data := *data
	if !ok {
		return
	}
	switch temp_data[key].(type) {
	case string:
		temp_data[key] = template.JSEscapeString(in)
		break

	case int:
		temp_data[key], ok = PostInt(key, c)
		if !ok {
			return
		}
		break

	case int32:
		temp_data[key], ok = PostInt64(key, c)
		if !ok {
			return
		}
		break

	case int64:
		temp_data[key], ok = PostInt64(key, c)
		if !ok {
			return
		}
		break

	case float64:
		temp_data[key], ok = PostFloat64(key, c)
		if !ok {
			return
		}
		break

	case float32:
		temp_data[key], ok = PostFloat64(key, c)
		if !ok {
			return
		}
		break

	case decimal.Decimal:
		temp_data[key], ok = PostDecimal(key, c)
		if !ok {
			return
		}
		break

	case nil:
		temp_data[key] = template.JSEscapeString(in)
		break

	case bool:
		temp_data[key], ok = PostBool(key, c)
		if !ok {
			return
		}
		break

	case time.Time:
		temp_data[key], ok = PostDateTime(key, c)
		if !ok {
			temp_data[key], ok = PostDate(key, c)
			if !ok {
				return
			}
		}
		break

	default:
		temp_data[key] = in
		break
	}
	data = &temp_data
	return true
}

func MPostDate(key string, c *gin.Context) (time.Time, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return time.Time{}, false
	} else {
		p, err := time.Parse("2006-01-02", in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should only be a Date"))
			c.Abort()
			return time.Time{}, false
		} else {
			return p, true
		}
	}
}

func MPostDateTime(key string, c *gin.Context) (time.Time, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return time.Time{}, false
	} else {
		p, err := time.Parse("2006-01-02 15:04:05", in)
		if err == nil {
			return p, true
		}
		p, err = time.Parse(time.RFC3339, in)
		if err == nil {
			return p, true
		}
		p, err = time.Parse(time.RFC3339Nano, in)
		if err == nil {
			return p, true
		}
		c.JSON(RET.Ret_fail(407, err.Error(), key+" should only be a DateTime or RFC3339"))
		c.Abort()
		return time.Time{}, false
	}
}

func MPostLength(key string, min, max int, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", false
	} else {
		err := Vali.Length(in, min, max)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" "+err.Error(), key+" "+err.Error()))
			c.Abort()
			return "", false
		}
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}
