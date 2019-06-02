package nagos

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCase(t *testing.T) {

	params := url.Values{}
	params["name"] = []string{"dick"}
	params["age"] = []string{"12"}

	encode := params.Encode()
	fmt.Println(encode)

}
