package Input

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/gorose-pro"
	"html/template"
	"main.go/tuuz/RET"
	"main.go/tuuz/Vali"
	"strings"
	"time"
)

func MPostAuto(c *gin.Context, goroseData *gorose.Data, where *map[string]interface{}) (ok bool, data map[string]interface{}) {
	whereMap := *where
	auto_keys := []string{}
	for key, _ := range *goroseData {
		_, whereHave := whereMap[key]
		if whereHave {
			auto_keys = append(auto_keys, key)
			okWhere, ret := MPost(key, c, goroseData)
			if !okWhere {
				c.JSON(RET.Ret_fail(400, nil, key+" should be exist or Not in the GoroseProWhere"))
				c.Abort()
				return false, nil
			}
			whereMap[key] = ret
		} else {
			auto_keys = append(auto_keys, key)
			okData, ret := MPost(key, c, goroseData)
			if okData {
				//if data's key is existed here then insert that data into the map, otherwise it won't shows in the datamap where it returns
				data[key] = ret
			}
		}
	}
	where = &whereMap
	if len(data) < 1 {
		c.JSON(RET.Ret_fail(400, "request in ["+strings.Join(auto_keys, ",")+"]", "GoroseProData is not ready"))
		c.Abort()
		return false, nil
	}
	return true, data
}

func MPostIn(c *gin.Context, goroseData *gorose.Data, data_keys []string) (ok bool, data map[string]interface{}) {
	temp_data := *goroseData
	data = make(map[string]interface{})
	for _, data_key := range data_keys {
		_, whereHave := temp_data[data_key]
		if whereHave {
			okWhere, ret := MPost(data_key, c, goroseData)
			if !okWhere {
				continue
			}
			data[data_key] = ret
		}
	}
	if len(data) < 1 {
		c.JSON(RET.Ret_fail(400, "request in ["+strings.Join(data_keys, ",")+"]", "GoroseProData is not ready"))
		c.Abort()
		return false, nil
	}
	return true, data
}

func MPost(key string, c *gin.Context, goroseData *gorose.Data) (ok bool, ret interface{}) {
	var in string
	in, ok = c.GetPostForm(key)
	if !ok {
		return
	}
	temp_data := *goroseData
	tdata, ok := temp_data[key]
	if !ok {
		return
	}
	switch tdata.(type) {
	case nil:
		ret = in
		return

	case int:
		ret, ok = PostInt(key, c)
		if !ok {
			return
		}
		break

	case int64:
		ret, ok = PostInt64(key, c)
		if !ok {
			ret, ok = PostBool(key, c)
			if !ok {
				return
			}
		}
		break

	case float64:
		ret, ok = PostFloat64(key, c)
		if !ok {
			return
		}
		break

	case time.Time:
		ret, ok = PostDateTime(key, c)
		if !ok {
			ret, ok = PostDate(key, c)
			if !ok {
				return
			}
		}
		break

	default:
		ret = template.JSEscapeString(in)
		break
	}
	temp_data[key] = ret
	goroseData = &temp_data
	return
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
