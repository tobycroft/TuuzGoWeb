package Input

import (
	"crypto/sha256"
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/tobycroft/gorose-pro"
	"html/template"
	"io"
	"main.go/config/app_conf"
	"main.go/tuuz/Array"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Date"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/RET"
	"main.go/tuuz/Vali"
	"path/filepath"
	"strings"
	"time"
)

func MPostAuto(data *gorose.Data, c *gin.Context, xss bool) {
	temp_data := *data
	for key, _ := range temp_data {
		p, ok := c.GetPostForm(key)
		if !ok {
			continue
		}
		temp_data[key] = p
	}
	data = &temp_data
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

	case int:
		temp_data[key], ok = PostInt(key, c)
		if !ok {
			return
		}

	case int32:
		temp_data[key], ok = PostInt64(key, c)
		if !ok {
			return
		}

	case int64:
		temp_data[key], ok = PostInt64(key, c)
		if !ok {
			return
		}

	case float64:
		temp_data[key], ok = PostFloat64(key, c)
		if !ok {
			return
		}

	case float32:
		temp_data[key], ok = PostFloat64(key, c)
		if !ok {
			return
		}

	case decimal.Decimal:
		temp_data[key], ok = SPostInt(key, c)
		if !ok {
			return
		}

	case nil:
		temp_data[key] = in

	case bool:
		temp_data[key], ok = PostBool(key, c)
		if !ok {
			return
		}

	}
	data = &temp_data
	return true
}

func MPostPhone(key string, length int, c *gin.Context) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", false
	} else {
		ret, err := decimal.NewFromString(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should only be numbers", key+" should only be numbers"))
			c.Abort()
			return "", false
		}
		err = Vali.Length(ret.String(), length, length)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" "+err.Error(), key+" "+err.Error()))
			c.Abort()
			return "", false
		}
		return ret.String(), true
	}
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

func MPostTime(key string, c *gin.Context) (time.Time, bool) {
	in, ok := PostInt64(key, c)
	if !ok {
		return time.Time{}, false
	} else {
		return time.Unix(in, 0), true
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

func MPostInt(key string, c *gin.Context) (int, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return 0, false
	} else {
		i, e := Calc.String2Int(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, e.Error(), key+" should be int"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func MPostInt64(key string, c *gin.Context) (int64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return 0, false
	} else {
		i, e := Calc.String2Int64(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, e.Error(), key+" should be int64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func MPostFloat64(key string, c *gin.Context) (float64, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return 0, false
	} else {
		i, e := Calc.String2Float64(in)
		if e != nil {
			c.JSON(RET.Ret_fail(407, e.Error(), key+" should be float64"))
			c.Abort()
			return 0, false
		}
		return i, true
	}
}

func MPostDecimal(key string, c *gin.Context) (decimal.Decimal, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return decimal.Zero, false
	} else {
		ret, err := decimal.NewFromString(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should be a Number"))
			c.Abort()
			return decimal.Zero, false
		}
		return ret, true
	}
}

func MPostBool(key string, c *gin.Context) (bool, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
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
			c.JSON(RET.Ret_fail(407, key+" should be Boolean", key+" should be Boolean"))
			c.Abort()
			return false, false
		}
	}
}

func MPostArray(key string, c *gin.Context) ([]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return nil, false
	} else {
		i, err := Jsong.JArray(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should be a Json-Array"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func MPostObject(key string, c *gin.Context) (map[string]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
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

func MPostArrayObject(key string, c *gin.Context) ([]map[string]interface{}, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
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

func MPostAny(key string, c *gin.Context, AnyType interface{}) bool {
	in, ok := c.GetPostForm(key)
	if !ok {
		return false
	} else {
		err := jsoniter.UnmarshalFromString(in, &AnyType)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should be a Json-AnyType"))
			c.Abort()
			return false
		}
		return true
	}
}

func MPostLimitPage(c *gin.Context) (int, int, error) {
	limit, ok := MPostInt("limit", c)
	if !ok {
		return 0, 0, errors.New("limit")
	}
	page, ok := MPostInt("page", c)
	if !ok {
		return 0, 0, errors.New("page")
	}
	return limit, page, nil
}

func MPostIn(key string, c *gin.Context, str_slices []string) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", false
	} else {
		if Array.InArrayString(in, str_slices) {
			return in, true
		} else {
			c.JSON(RET.Ret_fail(407, key+" 's data should in ["+strings.Join(str_slices, ",")+"]", key+" 's data should in ["+strings.Join(str_slices, ",")+"]"))
			c.Abort()
			return in, false
		}
	}
}

func MUpload(c *gin.Context) (File, bool) {
	file, err := c.FormFile("file")
	if err != nil {
		return File{}, false
	}
	temp_file, err := file.Open()
	defer temp_file.Close()
	if err != nil {
		c.JSON(RET.Ret_fail(406, "File Open Error", "POST-[file]:"+err.Error()))
		c.Abort()
		return File{}, false
	}
	file_hash := sha256.New()
	if _, err := io.Copy(file_hash, temp_file); err != nil {
		c.JSON(RET.Ret_fail(303, "File Hash Error", "POST-[file]:"+err.Error()))
		c.Abort()
		return File{}, false
	}
	file_md5 := file_hash.Sum(nil)
	filename := filepath.Base(file.Filename)
	ext := filepath.Ext(file.Filename)
	var path string
	if app_conf.FilePathCreateByDay {
		path = app_conf.FileSavePath + "/" + Date.TodayCombine() + "/"
	} else if app_conf.FilePathCreateByDate {
		path = app_conf.FileSavePath + "/" + Date.ThisMonthCombine() + "/"
	} else {
		path = app_conf.FileSavePath + "/"
	}
	err = pathmake(path)
	if err != nil {
		c.JSON(RET.Ret_fail(500, "Create Path Fail", "POST-[file]:"+err.Error()))
		c.Abort()
		return File{}, false
	}
	if app_conf.FileNameSecurity {
		filename = Calc.Md5(filename)
	}
	err = c.SaveUploadedFile(file, path+filename+ext)
	if err != nil {
		c.JSON(RET.Ret_fail(500, "File Saved Fail", "POST-[file]:"+err.Error()))
		c.Abort()
		return File{}, false
	}
	return File{
		Path:     path + "/" + filename,
		FileName: filename,
		Size:     file.Size,
		Md5:      string(file_md5),
		Mime:     file.Header.Get("Content-Type"),
		Ext:      ext,
	}, true
}
