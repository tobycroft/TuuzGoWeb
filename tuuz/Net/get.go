package Net

import (
	"crypto/tls"
)

type Get struct {
	curl               *Curl
	InsecureSkipVerify bool
	ret                *response
	err                error
}

func (self Get) Get(url string, queries map[string]interface{}, headers map[string]string, cookies map[string]string) Get {
	req := self.curl.NewRequest().request
	req.SetHeaders(headers)
	req.SetCookies(cookies)
	req.SetTimeout(5)
	req.DisableKeepAlives(true)
	req.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = req.Get(url, queries)
	return self
}

func (self Get) RetCookie() (cookie map[string]interface{}, err error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.curl.cookieHandler(self.ret.Cookies()), nil
}

func (self Get) RetString() (string, error) {
	if self.err != nil {
		return "", self.err
	}
	return self.ret.bodystring()
}

func (self Get) RetBytes() ([]byte, error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.ret.bodybytes()
}

func (self Get) RetJson(v any) error {
	if self.err != nil {
		return self.err
	}
	return self.ret.bodyjson(v)
}
