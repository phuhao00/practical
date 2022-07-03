package main

import (
	"fmt"
	"runtime"
)

//

func main() {
	runtime.GOMAXPROCS(3)
	SliceDemo()

}

func SliceDemo() {
	var s1 []int
	s2 := make([]int, 9)
	fmt.Println(s2)
	s1 = make([]int, 9, 9)
	print(s1)
	s1[0] = 0
	s1[1] = 1
	fmt.Println(s1[0:1])
	s1 = append(s1, 1) //
	println("---", s1[0])
	var m1 map[int]bool
	m1 = make(map[int]bool)
	m1[1] = true

	m1[5] = true
	m1[6] = true
	m1[2] = true
	m1[3] = true
	m1[4] = true

	for i, i2 := range s1 {
		fmt.Println(i, i2)
	}

	for i, b := range m1 {
		fmt.Println(i, b)
	}

}
