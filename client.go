package nagos

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// nacos客户端主入口
// config，service，都放这里了
type Client struct {
	RegistryConfigs []RegistryConfig

	// http客户端
	HttpClient HttpClient

	mux sync.Mutex

	Log log.Logger
}

type ClientConfig struct {
	RegistryConfig []RegistryConfig
}

func NewClient(clientConfig *ClientConfig) *Client {
	c := &Client{
		RegistryConfigs: clientConfig.RegistryConfig,
		HttpClient:      NewDefaultHttpClient(),
	}
	return c
}

func do(c *Client, path string, h func(url string) (*http.Response, error)) (string, error) {
	var msg string
	for i, rc := range c.RegistryConfigs {
		u := rc.UrlWithPath(path)

		// do something
		resp, err := h(u)

		if err != nil {
			if hasNextRegistry(i, &c.RegistryConfigs) {
				log.Printf("cannot connect to server %s, try next registry", u)
			} else {
				log.Printf("cannot connect to all server.")
				panic(err)
			}
			continue
		}

		code := resp.StatusCode
		if code != 200 {
			if code == 400 {
				log.Printf("get an error response from %s,bad request", u)
			} else if code == 403 {
				log.Printf("server %s,forbidden", u)
			} else if code == 404 {
				log.Printf("server %s,not found", u)
			} else if code == 500 {
				log.Printf("internal server error, server %s", u)
			} else {
				log.Printf("server %s response status code is %d", u, code)
			}
			log.Printf("response status code is %d", code)
			break
		}

		msg, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			panic(err)
		}

		return string(msg), nil
	}
	return "", errors.New("api request failed, response status code wasn't 200, msg: " + string(msg))
}

func hasNextRegistry(i int, r *[]RegistryConfig) bool {
	l := len(*r)
	return i < l
}
