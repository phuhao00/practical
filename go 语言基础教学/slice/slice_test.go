package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var s1 []int
	s1 = make([]int, 0)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s2 := s1[:2]
	fmt.Println(s2)
	s3 := s1[1:]
	fmt.Println(s3)
	s4 := s1[:1:2]
	fmt.Println(s4, cap(s4), len(s4))

	for index, val := range s1 {
		fmt.Println(index, val)
	}
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}

}
