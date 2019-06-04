package nagos

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	display := &CharDisplay{
		Char:'A',
	}

	result := display.Display(display)
	fmt.Println(result)
}
