package RET

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"strings"
)

var (
	jsonContentType      = []string{"application/json; charset=utf-8"}
	jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	jsonASCIIContentType = []string{"application/json"}
)

func json(c *gin.Context, retCodePointer *int, retJsonPointer any) {
	writeContentType(c.Writer, jsonContentType)
	body, err := sonic.MarshalString(&retJsonPointer)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.WriteString("RetData-Marshal-Error:" + err.Error())
		return
	}
	c.Writer.WriteHeader(*retCodePointer)
	c.Writer.WriteString(body)
}

func secure_json(c *gin.Context, retCodePointer *int, retJsonPointer any) {
	writeContentType(c.Writer, jsonContentType)
	body, err := sonic.MarshalString(&retJsonPointer)
	if err != nil {
		c.Writer.WriteHeader(500)
		c.Writer.WriteString("RetData-Marshal-Error:" + err.Error())
		return
	}

	if strings.HasPrefix(body, "[") && strings.HasSuffix(body, "]") {
		if _, err = c.Writer.WriteString(app_conf.SecureJsonPrefix); err != nil {
			return
		}
	}
	c.Writer.WriteHeader(*retCodePointer)
	c.Writer.WriteString(body)
}
