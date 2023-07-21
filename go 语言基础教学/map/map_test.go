package _map

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var m map[int]int
	m = make(map[int]int)
	m[1] = 1
	fmt.Println(m[1])
	for key, val := range m {
		fmt.Println(key, val)
	}
}
