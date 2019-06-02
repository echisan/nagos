package nagos

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewServerConfig(t *testing.T) {

	config := NewServerConfig("","","")
	config.SetTenant("tenant")

	bytes, _ := json.Marshal(config)
	fmt.Println(string(bytes))

}
