package 控制语句

import (
	"fmt"
	"testing"
)

// switch  if  else for  goto

func TestIfElse(t *testing.T) {
	if false {

	} else if false {

	} else {

	}
}

func TestSwitch(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	for _, i2 := range s1 {
		switch i2 {
		case 1:
			fmt.Println("A")
		case 2:
			fmt.Println("B")
		case 3:
			fmt.Println("C")
		case 4:
			fmt.Println("D")
		}
	}
}

func TestWhile(t *testing.T) {
	s2 := []int{1, 2, 3, 4, 5, 7}
	for i := 0; i < len(s2); i++ {
		fmt.Println(i, s2[i])
	}
}

func TestGoto(t *testing.T) {

}