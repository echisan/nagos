package nagos

import (
	"testing"
)

var rc = RegistryConfig{
	Host:        "localhost",
	Port:        8848,
	ContextPath: "/nacos",
}

var c = NewClient(&ClientConfig{RegistryConfig: []RegistryConfig{rc}})

func TestClient_PostConfig(t *testing.T) {

	cf := &Config{
		Content: "go-content",
		Group:   "go-group",
		DataId:  "go-data-id",
	}

	err := c.PostConfig(cf)
	if err != nil {
		panic(err)
	}

}
