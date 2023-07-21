package fn

import (
	"fmt"
	"testing"
)

func TestFn(t *testing.T) {

}

func A() {
	fmt.Println("A")
}
func AB() int {
	fmt.Println("AB")
	return 0
}

func ABC(ageIn int, nameIn string) (age int, name string) {
	fmt.Println(ageIn, nameIn)
	return ageIn, nameIn
}
