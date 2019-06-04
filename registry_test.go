package nagos

import (
	"testing"
)

func TestRegistryConfig_Url(t *testing.T) {
	c := &RegistryConfig{
		Host:        "127.0.0.1",
		Port:        8080,
		ContextPath: "/nacos",
	}
	if c.Url() != "http://127.0.0.1:8080/nacos" {
		t.Fatal("url failed")
	}
}

func TestRegistryConfig_UrlWithPath(t *testing.T) {
	c := &RegistryConfig{
		Host:        "127.0.0.1",
		Port:        8080,
		ContextPath: "/nacos",
	}
	path := c.UrlWithPath(PathConfigs)
	if path != "http://127.0.0.1:8080/nacos/v1/cs/configs" {
		t.Fatal("url with path failed")
	}
}
