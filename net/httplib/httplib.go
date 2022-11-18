package httplib

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

// var Log *zlog.Log

// func init() {
// 	Log = zlog.Init("20060102.http")
// }

var defaultSetting = HTTPSettings{
	ConnectTimeout:   60 * time.Second,
	ReadWriteTimeout: 60 * time.Second,
	Gzip:             true,
	DumpBody:         true,
}

// var defaultCookieJar http.CookieJar
// var settingMutex sync.Mutex

// createDefaultCookie creates a global cookiejar to store cookies.
// func createDefaultCookie() {
// 	settingMutex.Lock()
// 	defer settingMutex.Unlock()
// 	defaultCookieJar, _ = cookiejar.New(nil)
// }

type HTTPRequest struct {
	url       string
	req       *http.Request
	params    map[string][]string
	files     map[string]string
	setting   HTTPSettings
	resp      *http.Response
	body      []byte
	dump      []byte
	cookieJar http.CookieJar
}
type HTTPSettings struct {
	ShowDebug        bool
	UserAgent        string
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
	TLSClientConfig  *tls.Config
	Proxy            func(*http.Request) (*url.URL, error)
	Transport        http.RoundTripper
	CheckRedirect    func(req *http.Request, via []*http.Request) error
	EnableCookie     bool
	Gzip             bool
	DumpBody         bool
	Retries          int // if set to -1 means will retry forever
}

func NewRequest(rawurl, method string) *HTTPRequest {
	var resp http.Response
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Println("Httplib:", err)
	}
	req := http.Request{
		URL:        u,
		Method:     method,
		Header:     make(http.Header),
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
	}
	return &HTTPRequest{
		url:     rawurl,
		req:     &req,
		params:  map[string][]string{},
		files:   map[string]string{},
		setting: defaultSetting,
		resp:    &resp,
	}
}

func Get(url string) *HTTPRequest {
	return NewRequest(url, "GET")
}

func Post(url string) *HTTPRequest {
	return NewRequest(url, "POST")
}

func Put(url string) *HTTPRequest {
	return NewRequest(url, "PUT")
}

func Delete(url string) *HTTPRequest {
	return NewRequest(url, "DELETE")
}

func Head(url string) *HTTPRequest {
	return NewRequest(url, "HEAD")
}

func (b *HTTPRequest) Post(url string) *HTTPRequest {
	return b.NewRequest(url, "POST")
}

func (b *HTTPRequest) Get(url string) *HTTPRequest {
	return b.NewRequest(url, "GET")
}

func (b *HTTPRequest) Head(url string) *HTTPRequest {
	return b.NewRequest(url, "HEAD")
}

func (b *HTTPRequest) Delete(url string) *HTTPRequest {
	return b.NewRequest(url, "DELETE")
}

func (b *HTTPRequest) Put(url string) *HTTPRequest {
	return b.NewRequest(url, "PUT")
}

func (b *HTTPRequest) NewRequest(rawurl, method string) *HTTPRequest {
	var resp http.Response
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Println("Httplib:", err)
	}
	b.req.URL = u
	b.req.Method = method
	b.req.Header = make(http.Header)
	b.req.Proto = "HTTP/1.1"
	b.req.ProtoMajor = 1
	b.req.ProtoMinor = 1
	b.url = rawurl
	b.params = map[string][]string{}
	b.files = map[string]string{}
	b.resp = &resp
	return b
}

// Setting Change request settings
func (b *HTTPRequest) Setting(setting HTTPSettings) *HTTPRequest {
	b.setting = setting
	return b
}

// SetBasicAuth sets the request's Authorization header to use HTTP Basic Authentication with the provided username and password.
func (b *HTTPRequest) SetBasicAuth(username, password string) *HTTPRequest {
	b.req.SetBasicAuth(username, password)
	return b
}

// SetEnableCookie sets enable/disable cookiejar
func (b *HTTPRequest) SetEnableCookie(enable bool) *HTTPRequest {
	b.setting.EnableCookie = enable
	return b
}

// SetUserAgent sets User-Agent header field
func (b *HTTPRequest) SetUserAgent(useragent string) *HTTPRequest {
	b.setting.UserAgent = useragent
	return b
}

// Debug sets show debug or not when executing request.
func (b *HTTPRequest) Debug(isdebug bool) *HTTPRequest {
	b.setting.ShowDebug = isdebug
	return b
}

// SetTimeout sets connect time out and read-write time out for Request.
func (b *HTTPRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *HTTPRequest {
	b.setting.ConnectTimeout = connectTimeout
	b.setting.ReadWriteTimeout = readWriteTimeout
	return b
}

// SetTLSClientConfig sets tls connection configurations if visiting https url.
func (b *HTTPRequest) SetTLSClientConfig(config *tls.Config) *HTTPRequest {
	b.setting.TLSClientConfig = config
	return b
}

// Header add header item string in request.
func (b *HTTPRequest) Header(key, value string) *HTTPRequest {
	b.req.Header.Set(key, value)
	return b
}

// SetHost set the request host
func (b *HTTPRequest) SetHost(host string) *HTTPRequest {
	b.req.Host = host
	return b
}

// SetProtocolVersion Set the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *HTTPRequest) SetProtocolVersion(vers string) *HTTPRequest {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}
	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		b.req.Proto = vers
		b.req.ProtoMajor = major
		b.req.ProtoMinor = minor
	}
	return b
}

// SetCookie add cookie into request.
func (b *HTTPRequest) SetCookie(cookie *http.Cookie) *HTTPRequest {
	b.req.Header.Add("Cookie", cookie.String())
	return b
}

// SetTransport set the setting transport
func (b *HTTPRequest) SetTransport(transport http.RoundTripper) *HTTPRequest {
	b.setting.Transport = transport
	return b
}

// SetProxy set the http proxy
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (b *HTTPRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *HTTPRequest {
	b.setting.Proxy = proxy
	return b
}

// SetCheckRedirect specifies the policy for handling redirects.
//
// If CheckRedirect is nil, the Client uses its default policy,
// which is to stop after 10 consecutive requests.
func (b *HTTPRequest) SetCheckRedirect(redirect func(req *http.Request, via []*http.Request) error) *HTTPRequest {
	b.setting.CheckRedirect = redirect
	return b
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *HTTPRequest) Param(key, value string) *HTTPRequest {
	if param, ok := b.params[key]; ok {
		b.params[key] = append(param, value)
	} else {
		b.params[key] = []string{value}
	}
	return b
}

// PostFile add a post file to the request
func (b *HTTPRequest) PostFile(formname, filename string) *HTTPRequest {
	b.files[formname] = filename
	return b
}

// Body adds request raw body.
// it supports string and []byte.
func (b *HTTPRequest) Body(data interface{}) *HTTPRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	default:
		data1, _ := json.Marshal(data)
		bf := bytes.NewBuffer(data1)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(data1))
	}
	return b
}

// JSONBody adds request raw body encoding by JSON.
func (b *HTTPRequest) JSONBody(obj interface{}) (*HTTPRequest, error) {
	if b.req.Body == nil && obj != nil {
		byts, err := json.Marshal(obj)
		if err != nil {
			return b, err
		}
		b.req.Body = ioutil.NopCloser(bytes.NewReader(byts))
		b.req.ContentLength = int64(len(byts))
		b.req.Header.Set("Content-Type", "application/json")
	}
	return b, nil
}

func (b *HTTPRequest) buildURL(paramBody string) {
	// build GET url with query string
	if b.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Index(b.url, "?") != -1 {
			b.url += "&" + paramBody
		} else {
			b.url = b.url + "?" + paramBody
		}
		return
	}

	// build POST/PUT/PATCH url and body
	if (b.req.Method == "POST" || b.req.Method == "PUT" || b.req.Method == "PATCH") && b.req.Body == nil {
		// with files
		if len(b.files) > 0 {
			pr, pw := io.Pipe()
			bodyWriter := multipart.NewWriter(pw)
			go func() {
				for formname, filename := range b.files {
					fname := ""
					idx1 := strings.LastIndex(filename, `/`)
					idx2 := strings.LastIndex(filename, `\`)
					if idx1 == -1 && idx2 == -1 {
						fname = filename
					} else if idx1 > idx2 {
						fname = filename[idx1+1:]
					} else {
						fname = filename[idx2+1:]
					}
					fileWriter, err := bodyWriter.CreateFormFile(formname, fname)
					if err != nil {
						log.Println("Httplib:", err)
					}
					fh, err := os.Open(filename)
					if err != nil {
						log.Println("Httplib:", err)
					}
					//iocopy
					_, err = io.Copy(fileWriter, fh)
					fh.Close()
					if err != nil {
						log.Println("Httplib:", err)
					}
				}
				for k, v := range b.params {
					for _, vv := range v {
						bodyWriter.WriteField(k, vv)
					}
				}
				bodyWriter.Close()
				pw.Close()
			}()
			b.Header("Content-Type", bodyWriter.FormDataContentType())
			b.req.Body = ioutil.NopCloser(pr)
			return
		}

		// with params
		if len(paramBody) > 0 {
			b.Header("Content-Type", "application/x-www-form-urlencoded")
			b.Body(paramBody)
		}
	}
}

func (b *HTTPRequest) getResponse() (*http.Response, error) {
	if b.resp.StatusCode != 0 {
		return b.resp, nil
	}
	resp, err := b.DoRequest()
	if err != nil {
		return nil, err
	}
	b.resp = resp
	return resp, nil
}

// DoRequest will do the client.Do
func (b *HTTPRequest) DoRequest() (resp *http.Response, err error) {
	// var sTime, eTime time.Time
	// defer func() {
	// 	eTime = time.Now()
	// 	t := eTime.Sub(sTime).Nanoseconds() / 1000000
	// 	// data, _ := ioutil.ReadAll(resp.Body)
	// 	Log.Println(t, "ms", b.req.Method, "URL:", b.url, "DATA:", string(b.body), "RESPONSE:", resp.Status)
	// }()
	// sTime = time.Now()
	var paramBody string
	if len(b.params) > 0 {
		var buf bytes.Buffer
		for k, v := range b.params {
			for _, vv := range v {
				buf.WriteString(url.QueryEscape(k))
				buf.WriteByte('=')
				buf.WriteString(url.QueryEscape(vv))
				buf.WriteByte('&')
			}
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	b.buildURL(paramBody)
	url, err := url.Parse(b.url)
	if err != nil {
		return nil, err
	}

	b.req.URL = url

	trans := b.setting.Transport

	if trans == nil {
		// create default transport
		trans = &http.Transport{
			TLSClientConfig:     b.setting.TLSClientConfig,
			Proxy:               b.setting.Proxy,
			Dial:                TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout),
			MaxIdleConnsPerHost: -1,
		}
	} else {
		// if b.transport is *http.Transport then set the settings.
		if t, ok := trans.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = b.setting.TLSClientConfig
			}
			if t.Proxy == nil {
				t.Proxy = b.setting.Proxy
			}
			if t.Dial == nil {
				t.Dial = TimeoutDialer(b.setting.ConnectTimeout, b.setting.ReadWriteTimeout)
			}
		}
	}

	var jar http.CookieJar
	if b.setting.EnableCookie {
		// if defaultCookieJar == nil {
		// 	createDefaultCookie()
		// }
		// jar = defaultCookieJar
		if b.cookieJar == nil {
			b.cookieJar, _ = cookiejar.New(nil)
		}
		jar = b.cookieJar
	}

	client := &http.Client{
		Transport: trans,
		Jar:       jar,
	}

	if b.setting.UserAgent != "" && b.req.Header.Get("User-Agent") == "" {
		b.req.Header.Set("User-Agent", b.setting.UserAgent)
	}

	if b.setting.CheckRedirect != nil {
		client.CheckRedirect = b.setting.CheckRedirect
	}

	if b.setting.ShowDebug {
		dump, err := httputil.DumpRequest(b.req, b.setting.DumpBody)
		if err != nil {
			log.Println(err.Error())
		}
		b.dump = dump
	}
	// retries default value is 0, it will run once.
	// retries equal to -1, it will run forever until success
	// retries is setted, it will retries fixed times.
	for i := 0; b.setting.Retries == -1 || i <= b.setting.Retries; i++ {
		resp, err = client.Do(b.req)
		if err == nil {
			break
		}
	}
	return resp, err
}

// String returns the body string in response.
// it calls Response inner.
func (b *HTTPRequest) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *HTTPRequest) Bytes() ([]byte, error) {
	if b.body != nil {
		return b.body, nil
	}
	resp, err := b.getResponse()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	if b.setting.Gzip && resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		b.body, err = ioutil.ReadAll(reader)
	} else {
		b.body, err = ioutil.ReadAll(resp.Body)
	}
	return b.body, err
}

// ToFile saves the body data in response to one file.
// it calls Response inner.
func (b *HTTPRequest) ToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// b.req.Close = true
	resp, err := b.getResponse()
	if err != nil {
		return err
	}
	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// ToJSON returns the map that marshals from the body bytes as json in response .
// it calls Response inner.
func (b *HTTPRequest) ToJSON(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// ToXML returns the map that marshals from the body bytes as xml in response .
// it calls Response inner.
func (b *HTTPRequest) ToXML(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// Response executes request client gets response mannually.
func (b *HTTPRequest) Response() (*http.Response, error) {
	return b.getResponse()
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}
