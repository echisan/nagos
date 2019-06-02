package nagos

import "testing"

var rc RegistryConfig
var client Client

func init() {
	rc = RegistryConfig{
		Host:        "139.199.66.140",
		Port:        8848,
		ContextPath: "/nacos",
	}

	// "Content-Type",{"application/x-www-form-urlencoded"

	client = Client{
		RegistryConfigs: []RegistryConfig{rc},
		HttpClient:      NewDefaultHttpClient(),
	}
}

func TestClient_PostConfig(t *testing.T) {
	config := NewServerConfig("test-data-id", "test-group", "test-content")
	err := client.PostConfig(config)
	if err != nil {
		panic(err)
	}
}

func TestClient_GetConfig(t *testing.T) {
	err := client.GetConfig("test-data-id", "test-group", "")
	if err != nil {
		panic(err)
	}
}
