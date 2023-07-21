package ticker

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	tc := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-tc.C:
			fmt.Println("tc")
		}
	}
}
