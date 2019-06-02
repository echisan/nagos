package nagos

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

var DefaultClient = &http.Client{}

type HttpClient interface {
	Get(url string, header http.Header, values url.Values) (*http.Response, error)
	Post(url string, header http.Header, values url.Values) (*http.Response, error)
	Put(url string, header http.Header, values url.Values) (*http.Response, error)
	Delete(url string, header http.Header, values url.Values) (*http.Response, error)
}

type DefaultHttpClient struct {
	Header  http.Header
	Timeout time.Duration
}

func NewDefaultHttpClient() *DefaultHttpClient {

	client :=  &DefaultHttpClient{}

	header := make(map[string][]string)
	header["Content-Type"] = []string{"application/x-www-form-urlencoded"}

	client.Header = header
	client.Timeout = 3 * time.Second

	DefaultClient.Timeout = 3 * time.Second

	return client
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

	req, err := http.NewRequest(method, url, strings.NewReader(values.Encode()))
	req.Header = c.Header

	if header != nil {
		for k, v := range header {
			req.Header[k] = v
		}
	}

	if err != nil {
		return nil, err
	}

	resp, err := DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
