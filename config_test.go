package nagos

import (
	"fmt"
	"testing"
)

//var nc = NewClient(&ClientConfig{
//	RegistryConfig: []RegistryConfig{
//		{
//			Host:        "localhost",
//			Port:        8848,
//			ContextPath: contextPath,
//		},
//	},
//})

func TestClient_DelConfig(t *testing.T) {

	err := c.DelConfig(&ConfigParam{Group: "go-group", DataId: "go-data-id"})
	if err != nil {
		panic(err)
	}

}

func TestConfigListener(t *testing.T) {
	s2 := "2"
	fmt.Println(string(s2[0]))
}