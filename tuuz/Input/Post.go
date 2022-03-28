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
	"os"
	"path/filepath"
	"strings"
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

func PostNull(key string, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", true
	} else {
		if xss {
			return template.JSEscapeString(in), true
		} else {
			return in, true
		}
	}
}

func PostNullWithLength(key string, max_length int, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		return "", true
	} else {
		err := Vali.Length(in, 0, max_length)
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

func PostPhone(key string, length int, c *gin.Context) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
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
		return in, true
	}
}

func PostLength(key string, min, max int, c *gin.Context, xss bool) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return "", false
	} else {
		err := Vali.Length(in, min, max)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" "+err.Error(), key+" "+err.Error()))
			c.Abort()
			return "", false
		}
		return in, true
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

func PostDecimal(key string, c *gin.Context) (decimal.Decimal, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
		return decimal.Zero, false
	} else {
		ret, err := decimal.NewFromString(in)
		if err != nil {
			c.JSON(RET.Ret_fail(407, key+" should be a Number", key+" should be a Number"))
			c.Abort()
			return decimal.Zero, false
		}
		return ret, true
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

func PostLimitPage(c *gin.Context) (int, int, error) {
	limit, ok := PostInt("limit", c)
	if !ok {
		return 0, 0, errors.New("limit")
	}
	page, ok := PostInt("page", c)
	if !ok {
		return 0, 0, errors.New("page")
	}
	return limit, page, nil
}

func PostIn(key string, c *gin.Context, str_slices []string) (string, bool) {
	in, ok := c.GetPostForm(key)
	if !ok {
		c.JSON(RET.Ret_fail(400, key, "POST-["+key+"]"))
		c.Abort()
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

type File struct {
	Path     string
	FileName string
	Size     int64
	Md5      string
	Mime     string
	Ext      string
}

func Upload(c *gin.Context) (File, bool) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(RET.Ret_fail(400, "File Upload Error", "POST-[file]:"+err.Error()))
		c.Abort()
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

func pathmake(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
	return err
}
