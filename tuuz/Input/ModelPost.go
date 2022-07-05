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

/*
MPostAuto
if u wanna use this functions
1.load the data by find functions like single_object_data:=Api_find(id)
2.prepare the where map[string]interface{} the value can be prepared like nil or exact, for example where_maps:=map[string]interface{}{"key1":int64(0),"key2":nil}
3.if u set the value in the wheremap above, then if the request val is nor in the same type, it will echo failed-json to client

4.use ok,data_to_change:=MPostAuto(c,&single_object_data,&where_maps)
5.grab the data from this function
6.use update like Api_update(where_maps,data_to_change)
7.direct insert the data_to_change into GorosePro
8.for example: db.Where(where_maps) | db.Data(data_to_change)
*/
func MPostAuto(c *gin.Context, goroseProData *gorose.Data, where *map[string]interface{}) (ok bool, data map[string]interface{}) {
	ok = mpost_checker(c, goroseProData)
	if !ok {
		return false, nil
	}
	if *where == nil || where == nil {
		where = &map[string]interface{}{
			"id": int64(0),
		}
	}
	whereMap := *where
	auto_wheres := []string{}
	auto_datas := []string{}
	for key, _ := range *goroseProData {
		_, whereHave := whereMap[key]
		if whereHave {
			auto_wheres = append(auto_wheres, key)
			okWhere, ret := MPost(key, c, goroseProData)
			if !okWhere {
				c.JSON(RET.Ret_fail(400, nil, key+" should be exist or Not in the GoroseProWhere"))
				c.Abort()
				return false, nil
			}
			whereMap[key] = ret
		} else {
			auto_datas = append(auto_datas, key)
			okData, ret := MPost(key, c, goroseProData)
			if okData {
				//if data's key is existed here then insert that data into the map, otherwise it won't shows in the datamap where it returns
				data[key] = ret
			}
		}
	}
	where = &whereMap
	if len(data) < 1 {
		c.JSON(RET.Ret_fail(400, "request with ["+strings.Join(auto_wheres, ",")+"] request in ["+strings.Join(auto_datas, ",")+"]", "GoroseProData is not need to set"))
		c.Abort()
		return false, nil
	}
	return true, data
}

func mpost_checker(c *gin.Context, goroseProData *gorose.Data) bool {
	if *goroseProData == nil || goroseProData == nil {
		c.JSON(RET.Ret_fail(404, "GoroseProData is not ready", nil))
		c.Abort()
		return false
	}
	return true
}

/*
MPostIn
if u wanna use this functions
1.load the data by find functions like single_object_data:=Api_find(id)
2.use ok,data_to_change:=MPostIn(c,&single_object_data,[]stirng{"key1","key2"}
3.grab the data from this function
4.use update like Api_update(id,data_to_change)
5.direct insert the data_to_change into GorosePro
6.for example, db.Data(data_to_change)
*/
func MPostIn(c *gin.Context, goroseProData *gorose.Data, data_keys []string) (ok bool, data map[string]interface{}) {
	ok = mpost_checker(c, goroseProData)
	if !ok {
		return false, nil
	}
	if data_keys == nil || len(data_keys) < 1 {
		c.JSON(RET.Ret_fail(402, "DataKeys should be set", "GoroseProKeys is not ready"))
		c.Abort()
		return false, nil
	}
	temp_data := *goroseProData
	data = make(map[string]interface{})
	for _, data_key := range data_keys {
		_, whereHave := temp_data[data_key]
		if whereHave {
			okWhere, ret := MPost(data_key, c, goroseProData)
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

func MPost(key string, c *gin.Context, goroseProData *gorose.Data) (ok bool, ret interface{}) {
	ok = mpost_checker(c, goroseProData)
	if !ok {
		return false, nil
	}
	var in string
	in, ok = c.GetPostForm(key)
	if !ok {
		return
	}
	temp_data := *goroseProData
	tdata, ok := temp_data[key]
	if !ok {
		return
	}
	switch tdata.(type) {
	case nil:
		ret = in
		return

	case int:
		ret, ok = SPostInt(key, c)
		if !ok {
			return
		}
		break

	case int64:
		ret, ok = SPostInt64(key, c)
		if !ok {
			ret, ok = SPostBool(key, c)
			if !ok {
				return
			}
		}
		break

	case float64:
		ret, ok = SPostFloat64(key, c)
		if !ok {
			return
		}
		break

	case time.Time:
		ret, ok = SPostDateTime(key, c)
		if !ok {
			ret, ok = SPostDate(key, c)
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
	goroseProData = &temp_data
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
