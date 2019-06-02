package nagos

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	RegistryConfigs []RegistryConfig

	HttpClient HttpClient
}

func (c *Client) PostConfig(config *Config) error {

	for _, rc := range c.RegistryConfigs {

		u := fmt.Sprintf("%s://%s:%d%s%s", protocol, rc.Host, rc.Port, rc.ContextPath, "/v1/cs/configs")

		fmt.Println(u)

		values := url.Values{"dataId": {config.DataId}, "group": {config.Group}, "content": {config.Content}}

		if config.Tenant != "" {
			values.Set("tenant", config.Tenant)
		}

		resp, err := c.HttpClient.Post(u, nil, values)

		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		msg := string(bytes)
		if msg == "true" {
			break
		} else {
			log.Fatalf("post config failed, msg:%s", msg)
			time.Sleep(3 * time.Second)
		}
	}
	return nil
}

func (c *Client) GetConfig(dataId, group, tenant string) error {
	for _, rc := range c.RegistryConfigs {
		u := fmt.Sprintf("%s://%s:%d%s%s", protocol, rc.Host, rc.Port, rc.ContextPath, "/v1/cs/configs")

		values := url.Values{}
		values.Set("dataId", dataId)
		values.Set("group", group)
		if tenant != "" {
			values.Set("tenant", tenant)
		}

		v := http.Header{}
		v.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJuYWNvcyIsImF1dGgiOiIiLCJleHAiOjE1NTkzMTQzMDZ9.3D0sI9IRNAowIDF81Ibk3VJ49kO2GeBb4AnDCtQX3GY")
		resp, err := c.HttpClient.Get(u+"/"+values.Encode(), v, nil)
		if err != nil {
			return err
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		msg := string(bytes)

		fmt.Println(msg)

		//if msg == "true" {
		//	break
		//} else {
		//	log.Fatalf("post config failed, msg:%s", msg)
		//	time.Sleep(3 * time.Second)
		//}
		return nil
	}
	return nil
}
