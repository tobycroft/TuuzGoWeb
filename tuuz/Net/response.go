package Net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

type response struct {
	time int64
	url  string
	resp *http.Response
	body []byte
}

func (r *response) Response() *http.Response {
	if r != nil {
		return r.resp
	}
	return nil
}

func (r *response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode
}

func (r *response) Time() string {
	if r != nil {
		return fmt.Sprintf("%dms", r.time)
	}
	return "0ms"
}

func (r *response) Url() string {
	if r != nil {
		return r.url
	}
	return ""
}

func (r *response) Headers() http.Header {
	if r != nil {
		return r.resp.Header
	}
	return nil
}

func (r *response) Cookies() []*http.Cookie {
	if r != nil {
		return r.resp.Cookies()
	}
	return []*http.Cookie{}
}

func (r *response) bodybytes() ([]byte, error) {
	if r == nil {
		return []byte{}, errors.New("HttpRequest.response is nil.")
	}

	defer r.resp.Body.Close()

	if len(r.body) > 0 {
		return r.body, nil
	}

	if r.resp == nil || r.resp.Body == nil {
		return nil, errors.New("response or body is nil")
	}

	b, err := io.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}
	r.body = b

	return b, nil
}

func (r *response) bodystring() (string, error) {
	b, err := r.bodybytes()
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (r *response) bodyjson(v interface{}) error {
	b, err := r.bodybytes()
	if err != nil {
		return err
	}
	if err := sonic.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

func (r *response) Close() error {
	if r != nil {
		return r.resp.Body.Close()
	}
	return nil
}

func (r *response) JsonInPrettify() (ret string, err error) {
	b, err := r.bodybytes()
	if err != nil {
		return "", err
	}
	if sonic.Valid(b) {
		return "", errors.New("illegal json: " + err.Error())
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return
	}
	ret = buf.String()
	return
}
