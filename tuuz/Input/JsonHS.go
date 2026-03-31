package Input

import (
	"bytes"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin/binding"
)

var JsonHS = sonicJSONBinding{}

type sonicJSONBinding struct{}

func (sonicJSONBinding) Name() string {
	return "json"
}

func (sonicJSONBinding) Bind(req *http.Request, obj any) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	// 恢复 body（避免后续读取失败）
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := sonic.Unmarshal(body, obj); err != nil {
		return err
	}

	// 复用 Gin 的 validator
	return binding.Validator.ValidateStruct(obj)
}
