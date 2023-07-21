package channel

import "testing"

func TestCh(t *testing.T) {

	var c1 chan int
	c1 = make(chan int)
	var c2 <-chan int
	c2 = make(chan int)
	var c3 chan<- int
	c3 = make(chan int)
	c4 := make(chan int, 3)
	fn1(c1)
	fn2(c2)
	fn3(c3)
	fn4(c4)
}

func fn1(chan int) {

}

func fn2(<-chan int) {

}

func fn3(chan<- int) {

}

func fn4(chan int) {

}

