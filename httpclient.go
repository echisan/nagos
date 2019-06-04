package nagos

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

const DefaultTimeout = 3 * time.Second
const DefaultContentType = "application/x-www-form-urlencoded"

type HttpClient interface {
	Get(url string, header http.Header, values url.Values) (*http.Response, error)
	Post(url string, header http.Header, values url.Values) (*http.Response, error)
	Put(url string, header http.Header, values url.Values) (*http.Response, error)
	Delete(url string, header http.Header, values url.Values) (*http.Response, error)
}

type DefaultHttpClient struct {
	Timeout time.Duration
	client  http.Client
}

func (c *DefaultHttpClient) Get(url string, header http.Header, values url.Values) (*http.Response, error) {
	return c.doRequest(http.MethodGet, url, header, values)
}

func (c *DefaultHttpClient) Post(url string, header http.Header, values url.Values) (*http.Response, error) {
	return c.doRequest(http.MethodPost, url, header, values)
}

func (c *DefaultHttpClient) Put(url string, header http.Header, values url.Values) (*http.Response, error) {
	return c.doRequest(http.MethodPut, url, header, values)
}

func (c *DefaultHttpClient) Delete(url string, header http.Header, values url.Values) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, url, header, values)
}

func (c *DefaultHttpClient) doRequest(method string, url string, header http.Header, values url.Values) (*http.Response, error) {

	var req *http.Request
	var err error

	if method == http.MethodGet {
		req, err = c.NewRequest(method, url+"?"+values.Encode(), header, nil)
	} else {
		req, err = c.NewRequest(method, url, header, values)
	}

	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewDefaultHttpClient() *DefaultHttpClient {
	d := DefaultHttpClient{}
	d.Timeout = DefaultTimeout
	d.client = http.Client{}
	return &d
}

func (*DefaultHttpClient) NewRequest(method string, url string, header http.Header, values url.Values) (*http.Request, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	// set header
	req.Header.Set("Content-Type", DefaultContentType)
	if header != nil {
		for k, v := range header {
			for _, s := range v {
				header.Add(k, s)
			}
		}
	}
	return req, nil
}
