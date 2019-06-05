package nagos

import (
	"errors"
	"net/http"
	"net/url"
)

type ConfigListener struct {
	DataId string
	Group string
	listener func(received string)
}

var cfls = &Configs{}

type Configs struct {
	listeners []*ConfigListener
}

func (c *Client) DelConfig(param *ConfigParam) error {
	if param.DataId == "" || param.Group == "" {
		panic("dataId or group cannot be empty")
	}
	t := param.Tenant

	v := url.Values{}
	v.Set("dataId", param.DataId)
	v.Set("group", param.Group)
	if t != "" {
		v.Set("tenant", t)
	}

	s, e := do(c, PathConfigs, func(url string) (response *http.Response, e error) {
		return c.HttpClient.Delete(url, nil, v)
	})

	if e != nil {
		return e
	}

	if s != "true" {
		return errors.New("delete config failed")
	}
	return nil
}

func (c *Client) PostConfig(config *Config) error {

	v := url.Values{}
	v.Set("dataId", config.DataId)
	v.Set("group", config.Group)
	v.Set("content", config.Content)
	if config.Tenant != "" {
		v.Set("tenant", config.Tenant)
	}

	s, e := do(c, PathConfigs, func(url string) (response *http.Response, e error) {
		return c.HttpClient.Post(url, nil, v)
	})

	if e != nil {
		return e
	}

	if s != "true" {
		return errors.New("post config failed")
	}

	return nil
}

type ConfigParam struct {
	DataId string
	Group  string
	Tenant string
}

func (c *Client) GetConfig(param *ConfigParam) (string, error) {
	if param.DataId == "" || param.Group == "" {
		panic("dataId or group cannot be empty")
	}
	t := param.Tenant

	v := url.Values{}
	v.Set("dataId", param.DataId)
	v.Set("group", param.Group)
	if t != "" {
		v.Set("tenant", t)
	}

	str, err := do(c, PathConfigs, func(url string) (response *http.Response, e error) {
		return c.HttpClient.Get(url, nil, v)
	})

	if err != nil {
		return "", err
	}

	return str, nil
}

// add configListener
func (c *Client) AddConfigListener(cfl *ConfigListener) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cfls.listeners = append(cfls.listeners, cfl)
}

func listen(cfl *ConfigListener) {
}

