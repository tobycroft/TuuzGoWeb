package Net

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func (r *Curl) NewRequest() *Curl {
	req := request{}
	req.SetTimeout(30)
	req.SetHeaders(map[string]string{})
	req.SetCookies(map[string]string{})
	req.Transport(transport)
	r.request = req
	return r
}

type request struct {
	cli               *http.Client
	transport         *http.Transport
	debug             bool
	url               string
	method            string
	time              int64
	timeout           time.Duration
	headers           map[string]string
	cookies           map[string]string
	username          string
	password          string
	data              interface{}
	disableKeepAlives bool
	tlsClientConfig   *tls.Config
	jar               http.CookieJar
	proxy             func(*http.Request) (*url.URL, error)
	checkRedirect     func(req *http.Request, via []*http.Request) error
}

func (r *request) DisableKeepAlives(v bool) *request {
	r.disableKeepAlives = v
	return r
}

func (r *request) Jar(v http.CookieJar) *request {
	r.jar = v
	return r
}

func (r *request) CheckRedirect(v func(req *http.Request, via []*http.Request) error) *request {
	r.checkRedirect = v
	return r
}

func (r *request) TLSClient(v *tls.Config) *request {
	return r.SetTLSClient(v)
}

func (r *request) SetTLSClient(v *tls.Config) *request {
	r.tlsClientConfig = v
	return r
}

func (r *request) Proxy(v func(*http.Request) (*url.URL, error)) *request {
	r.proxy = v
	return r
}

func (r *request) Transport(v *http.Transport) *request {
	r.transport = v
	return r
}

// Debug model
func (r *request) Debug(v bool) *request {
	r.debug = v
	return r
}

// Get transport
func (r *request) getTransport() http.RoundTripper {
	if r.transport == nil {
		return http.DefaultTransport
	}

	r.transport.DisableKeepAlives = r.disableKeepAlives

	if r.tlsClientConfig != nil {
		r.transport.TLSClientConfig = r.tlsClientConfig
	}

	if r.proxy != nil {
		r.transport.Proxy = r.proxy
	}

	return http.RoundTripper(r.transport)
}

// Build client
func (r *request) buildClient() *http.Client {
	if r.cli == nil {
		r.cli = &http.Client{
			Transport:     r.getTransport(),
			Jar:           r.jar,
			CheckRedirect: r.checkRedirect,
			Timeout:       time.Second * r.timeout,
		}
	}
	return r.cli
}

// Set headers
func (r *request) SetHeaders(headers map[string]string) *request {
	if headers != nil || len(headers) > 0 {
		for k, v := range headers {
			r.headers[k] = v
		}
	}
	return r
}

// Init headers
func (r *request) initHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range r.headers {
		req.Header.Set(k, v)
	}
}

// Set cookies
func (r *request) SetCookies(cookies map[string]string) *request {
	if cookies != nil || len(cookies) > 0 {
		for k, v := range cookies {
			r.cookies[k] = v
		}
	}
	return r
}

// Init cookies
func (r *request) initCookies(req *http.Request) {
	for k, v := range r.cookies {
		req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

// Set basic auth
func (r *request) SetBasicAuth(username, password string) *request {
	r.username = username
	r.password = password
	return r
}

func (r *request) initBasicAuth(req *http.Request) {
	if r.username != "" && r.password != "" {
		req.SetBasicAuth(r.username, r.password)
	}
}

// Check application/json
func (r *request) isJson() bool {
	if len(r.headers) > 0 {
		for _, v := range r.headers {
			if strings.Contains(strings.ToLower(v), "application/json") {
				return true
			}
		}
	}
	return false
}

// Build query data
func (r *request) buildBody(d ...interface{}) (io.Reader, error) {
	// GET and DELETE request dose not send body
	if r.method == "GET" || r.method == "DELETE" {
		return nil, nil
	}

	if len(d) == 0 || d[0] == nil {
		return strings.NewReader(""), nil
	}

	switch d[0].(type) {
	case string:
		return strings.NewReader(d[0].(string)), nil

	case map[string]any:
		break

	default:
		return strings.NewReader(""), errors.New("incorrect parameter format.")

	}
	if r.isJson() {
		if b, err := sonic.Marshal(d[0]); err != nil {
			return nil, err
		} else {
			return bytes.NewReader(b), nil
		}
	}

	data := make([]string, 0)
	for k, v := range d[0].(map[string]interface{}) {
		if s, ok := v.(string); ok {
			data = append(data, fmt.Sprintf("%s=%v", k, s))
			continue
		}
		b, err := sonic.Marshal(v)
		if err != nil {
			return nil, err
		}
		data = append(data, fmt.Sprintf("%s=%s", k, string(b)))
	}

	return strings.NewReader(strings.Join(data, "&")), nil
}

func (r *request) SetTimeout(d time.Duration) *request {
	r.timeout = d
	return r
}

// Parse query for GET request
func parseQuery(url string) ([]string, error) {
	urlList := strings.Split(url, "?")
	if len(urlList) < 2 {
		return make([]string, 0), nil
	}
	query := make([]string, 0)
	for _, val := range strings.Split(urlList[1], "&") {
		v := strings.Split(val, "=")
		if len(v) < 2 {
			return make([]string, 0), errors.New("query parameter error")
		}
		query = append(query, fmt.Sprintf("%s=%s", v[0], v[1]))
	}
	return query, nil
}

// Build GET request url
func buildUrl(url string, data ...interface{}) (string, error) {
	query, err := parseQuery(url)
	if err != nil {
		return url, err
	}
	if len(data) > 0 && data[0] != nil {
		switch data[0].(type) {
		case map[string]any:
			for k, v := range data[0].(map[string]any) {
				query = append(query, fmt.Sprintf("%s=%s", k, v))
			}
			break
		case map[string]string:
			for k, v := range data[0].(map[string]string) {
				query = append(query, fmt.Sprintf("%s=%s", k, v))
			}
			break
		case string:
			param := data[0].(string)
			if param != "" {
				query = append(query, param)
			}
			break
		default:
			return url, errors.New("incorrect parameter format.")
		}

	}
	list := strings.Split(url, "?")
	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", list[0], strings.Join(query, "&")), nil
	}
	return list[0], nil
}

func (r *request) elapsedTime(n int64, resp *response) {
	end := time.Now().UnixNano() / 1e6
	resp.time = end - n
}

func (r *request) log() {
	if r.debug {
		fmt.Printf("[HttpRequest]\n")
		fmt.Printf("-------------------------------------------------------------------\n")
		fmt.Printf("request: %s %s\nHeaders: %v\nCookies: %v\nTimeout: %ds\nReqBody: %v\n\n", r.method, r.url, r.headers, r.cookies, r.timeout, r.data)
		//fmt.Printf("-------------------------------------------------------------------\n\n")
	}
}

// Get is a get http request
func (r *request) Get(url string, data ...interface{}) (*response, error) {
	return r.request(http.MethodGet, url, data...)
}

// Post is a post http request
func (r *request) post(url string, data ...interface{}) (*response, error) {
	return r.request(http.MethodPost, url, data...)
}

// Put is a put http request
func (r *request) Put(url string, data ...interface{}) (*response, error) {
	return r.request(http.MethodPut, url, data...)
}

// Delete is a delete http request
func (r *request) Delete(url string, data ...interface{}) (*response, error) {
	return r.request(http.MethodDelete, url, data...)
}

// Upload file
func (r *request) Upload(url, filename, fileinput string) (*response, error) {
	return r.sendFile(url, filename, fileinput)
}

// Send http request
func (r *request) request(method, url string, data ...interface{}) (*response, error) {
	// Build response
	response := &response{}

	// Start time
	start := time.Now().UnixNano() / 1e6
	// Count elapsed time
	defer r.elapsedTime(start, response)

	if method == "" || url == "" {
		return nil, errors.New("parameter method and url is required")
	}

	// Debug infomation
	defer r.log()

	r.url = url
	if len(data) > 0 {
		r.data = data[0]
	} else {
		r.data = ""
	}

	var (
		err  error
		req  *http.Request
		body io.Reader
	)
	r.cli = r.buildClient()

	method = strings.ToUpper(method)
	r.method = method

	if method == "GET" || method == "DELETE" {
		url, err = buildUrl(url, data...)
		if err != nil {
			return nil, err
		}
		r.url = url
	}

	body, err = r.buildBody(data...)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	r.initHeaders(req)
	r.initCookies(req)
	r.initBasicAuth(req)

	resp, err := r.cli.Do(req)
	if err != nil {
		return nil, err
	}

	response.url = url
	response.resp = resp

	return response, nil
}

// Send file
func (r *request) sendFile(url, filename, fileinput string) (*response, error) {
	if url == "" {
		return nil, errors.New("parameter url is required")
	}

	fileBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(fileBuffer)
	fileWriter, er := bodyWriter.CreateFormFile(fileinput, filename)
	if er != nil {
		return nil, er
	}

	f, er := os.Open(filename)
	if er != nil {
		return nil, er
	}
	defer f.Close()

	_, er = io.Copy(fileWriter, f)
	if er != nil {
		return nil, er
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// Build response
	response := &response{}

	// Start time
	start := time.Now().UnixNano() / 1e6
	// Count elapsed time
	defer r.elapsedTime(start, response)

	// Debug infomation
	defer r.log()

	r.url = url
	r.data = nil

	var (
		err error
		req *http.Request
	)
	r.cli = r.buildClient()
	r.method = "POST"

	req, err = http.NewRequest(r.method, url, fileBuffer)
	if err != nil {
		return nil, err
	}

	r.initHeaders(req)
	r.initCookies(req)
	r.initBasicAuth(req)
	req.Header.Set("Content-Type", contentType)

	resp, err := r.cli.Do(req)
	if err != nil {
		return nil, err
	}

	response.url = url
	response.resp = resp

	return response, nil
}
