package RET

import (
	"bytes"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
)

var (
	jsonContentType      = []string{"application/json; charset=utf-8"}
	jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	jsonASCIIContentType = []string{"application/json"}
)

func json(c *gin.Context, retCode *int, retJsonPointer any) {
	writeContentType(c.Writer, jsonContentType)
	jsonBytes, err := jsoniter.Marshal(&retJsonPointer)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.WriteString("RetData-Marshal-Error:" + err.Error())
		return
	}
	c.Writer.WriteHeader(*retCode)
	c.Writer.Write(jsonBytes)
}

func secure_json(c *gin.Context, retCode *int, retJsonPointer any) {
	writeContentType(c.Writer, jsonContentType)
	jsonBytes, err := jsoniter.Marshal(&retJsonPointer)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.WriteString("RetData-Marshal-Error:" + err.Error())
		return
	}
	if bytes.HasPrefix(jsonBytes, StringToBytes("[")) && bytes.HasSuffix(jsonBytes,
		StringToBytes("]")) {
		if _, err = c.Writer.Write(StringToBytes(app_conf.SecureJsonPrefix)); err != nil {
			return
		}
	}
	c.Writer.WriteHeader(*retCode)
	c.Writer.Write(jsonBytes)
}
