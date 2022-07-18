package Input

import (
	"crypto/sha256"
	"errors"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
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

func SPost(key string, c *gin.Context, DemoType interface{}) interface{} {
	in, ok := c.GetPostForm(key)
	if !ok {
		return nil
	} else {
		switch DemoType.(type) {
		case string:
			return in

		case int:
			str, err := Calc.String2Int(in)
			if err != nil {
				return nil
			}
			return str

		case int32:
			str, err := Calc.String2Int64(in)
			if err != nil {
				return nil
			}
			return str

		case int64:
			str, err := Calc.String2Int64(in)
			if err != nil {
				return nil
			}
			return str

		case float64:
			str, err := Calc.String2Float64(in)
			if err != nil {
				return nil
			}
			return str

		case float32:
			str, err := Calc.String2Float64(in)
			if err != nil {
				return nil
			}
			return str

		case decimal.Decimal:
			ret, err := decimal.NewFromString(in)
			if err != nil {
				return nil
			}
			return ret

		case nil:
			return nil

		case bool:
			str, ok := SPostBool(key, c)
			if !ok {
				return nil
			}
			return str

		}
		return nil
	}
}

func SPostDefault[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal | any](key string, c *gin.Context, defaultValue T) T {
	in, ok := c.GetPostForm(key)
	if !ok {
		return defaultValue
	} else {
		switch any(defaultValue).(type) {
		case string:
			return any(in).(T)

		case int:
			str, err := Calc.String2Int(in)
			if err != nil {
				return defaultValue
			}
			return any(str).(T)

		case int32:
			str, err := Calc.String2Int64(in)
			if err != nil {
				return defaultValue
			}
			return any(str).(T)

		case int64:
			str, err := Calc.String2Int64(in)
			if err != nil {
				return defaultValue
			}
			return any(str).(T)

		case float64:
			str, err := Calc.String2Float64(in)
			if err != nil {
				return defaultValue
			}
			return any(str).(T)

		case float32:
			str, err := Calc.String2Float64(in)
			if err != nil {
				return defaultValue
			}
			return any(str).(T)

		case decimal.Decimal:
			ret, err := decimal.NewFromString(in)
			if err != nil {
				return defaultValue
			}
			return any(ret).(T)

		case nil:
			return defaultValue

		case bool:
			str, ok := SPostBool(key, c)
			if !ok {
				return defaultValue
			}
			return any(str).(T)

		}
		return defaultValue
	}
}

func SPostString(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", false
	} else {
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}

func SPostPhone(key string, length int, c *gin.Context) (string, bool) {
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

func SPostDate(key string, c *gin.Context) (time.Time, bool) {
	return SPostDateTime(key, c)
}

func SPostDateTime(key string, c *gin.Context) (time.Time, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return time.Time{}, false
	} else {
		datetime, err := Date.Date_time_parser(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should only be a Date(+Time) or RFC3339"))
			c.Abort()
			return time.Time{}, false
		}
		return datetime, true
	}
}

func SPostTime(key string, c *gin.Context) (time.Time, bool) {
	in, ok := PostInt64(key, c)
	if !ok {
		return time.Time{}, false
	} else {
		return time.Unix(in, 0), true
	}
}

func SPostLength(key string, min, max int, c *gin.Context, xss bool) (string, bool) {
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

func SPostInt(key string, c *gin.Context) (int, bool) {
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

func SPostInt64(key string, c *gin.Context) (int64, bool) {
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

func SPostFloat64(key string, c *gin.Context) (float64, bool) {
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

func SPostDecimal(key string, c *gin.Context) (decimal.Decimal, bool) {
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

func SPostBool(key string, c *gin.Context) (bool, bool) {
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

func SPostArray[T string | int | int32 | int64 | float32 | float64 | any](key string, c *gin.Context) ([]T, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return nil, false
	} else {
		i, err := Jsong.JArray[T](in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, err.Error(), key+" should be a Json-Array"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func SPostObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](key string, c *gin.Context) (map[T]V, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return nil, false
	} else {
		i, err := Jsong.JObject[T, V](in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-Object", key+" should be a Json-Object"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func SPostArrayObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](key string, c *gin.Context) ([]map[T]V, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return nil, false
	} else {
		i, err := Jsong.JArrayObject[T, V](in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Json-ArrayObject", key+" should be a Json-ArrayObject"))
			c.Abort()
			return nil, false
		}
		return i, true
	}
}

func SPostAny(key string, c *gin.Context, AnyType interface{}) bool {
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

func SPostLimitPage(c *gin.Context) (int, int, error) {
	limit, ok := SPostInt("limit", c)
	if !ok {
		return 0, 0, errors.New("limit")
	}
	page, ok := SPostInt("page", c)
	if !ok {
		return 0, 0, errors.New("page")
	}
	return limit, page, nil
}

func SPostIn(key string, c *gin.Context, str_slices []string) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", false
	} else {
		if Array.InArray(in, str_slices) {
			return in, true
		} else {
			c.JSON(RET.Ret_fail(407, key+" 's data should in ["+strings.Join(str_slices, ",")+"]", key+" 's data should in ["+strings.Join(str_slices, ",")+"]"))
			c.Abort()
			return in, false
		}
	}
}

func SUpload(c *gin.Context) (File, bool) {
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
