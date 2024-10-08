package Net

import (
	"crypto/tls"
	"time"
)

type Net struct {
	curl               Curl
	InsecureSkipVerify bool
	ret                *response
	err                error
	url                string
	query              map[string]string
	postData           map[string]string
	header             map[string]string
	cookie             map[string]string
}

func (self Net) New() *Net {
	self.curl.newRequest()
	return &self
}

func (self *Net) SetTimeOut(Timeout time.Duration) *Net {
	self.curl.request.SetTimeout(Timeout)
	return self
}
func (self *Net) AllowInsecure() *Net {
	self.InsecureSkipVerify = true
	return self
}

func (self *Net) SetUrl(url string) *Net {
	self.url = url
	return self
}

func (self *Net) SetQuery(query map[string]string) *Net {
	self.query = query
	return self
}

func (self *Net) SetPostData(postData map[string]string) *Net {
	self.postData = postData
	return self
}

func (self *Net) SetHeader(header map[string]string) *Net {
	self.header = header
	return self
}

func (self *Net) SetCookies(cookies map[string]string) *Net {
	self.cookie = cookies
	return self
}

func (self *Net) SetBasicAuth(username, password string) *Net {
	self.curl.request.SetBasicAuth(username, password)
	return self
}

func (self *Net) NetRpc() *Net {
	self.curl.SetHeaderJson()
	self.curl.request.DisableKeepAlives(true)
	self.curl.request.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = self.curl.request.post(self.url, self.postData)
	return self
}

func (self *Net) PostRaw() *Net {
	self.curl.SetHeaderTextPlain()
	self.curl.request.DisableKeepAlives(true)
	self.curl.request.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.ret, self.err = self.curl.request.post(self.url, self.postData)
	return self
}

func (self *Net) PostFormData() *Net {
	self.curl.SetHeaderFormData()
	self.curl.request.SetHeaders(self.header)
	self.curl.request.SetCookies(self.cookie)
	self.curl.request.DisableKeepAlives(true)
	self.curl.request.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.url, self.err = buildUrl(self.url, self.query)
	if self.err != nil {
		return self
	}
	self.ret, self.err = self.curl.request.postFD(self.url, self.postData)
	return self
}

func (self *Net) PostUrlXEncode() *Net {
	self.curl.SetHeaderUrlEncode()
	self.curl.request.SetHeaders(self.header)
	self.curl.request.SetCookies(self.cookie)
	self.curl.request.DisableKeepAlives(true)
	self.curl.request.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.url, self.err = buildUrl(self.url, self.query)
	if self.err != nil {
		return self
	}
	self.ret, self.err = self.curl.request.post(self.url, self.postData)
	return self
}

func (self *Net) PostJson() *Net {
	self.curl.SetHeaderJson()
	self.curl.request.SetHeaders(self.header)
	self.curl.request.SetCookies(self.cookie)
	self.curl.request.DisableKeepAlives(true)
	self.curl.request.SetTLSClient(&tls.Config{InsecureSkipVerify: self.InsecureSkipVerify})
	self.url, self.err = buildUrl(self.url, self.query)
	if self.err != nil {
		return self
	}
	self.ret, self.err = self.curl.request.post(self.url, self.postData)
	return self
}

func (self *Net) RetCookie() (cookie map[string]interface{}, err error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.curl.cookieHandler(self.ret.Cookies()), nil
}

func (self *Net) RetString() (string, error) {
	if self.err != nil {
		return "", self.err
	}
	return self.ret.bodystring()
}

func (self *Net) RetBytes() ([]byte, error) {
	if self.err != nil {
		return nil, self.err
	}
	return self.ret.bodybytes()
}

func (self *Net) RetJson(v any) error {
	if self.err != nil {
		return self.err
	}
	return self.ret.bodyjson(v)
}
