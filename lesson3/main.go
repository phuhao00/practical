package main

import (
	"fmt"
	"sync"
)

var w2 sync.WaitGroup

func main() {
	w2.Add(2)
	ChDemo()
	ChDemo2()
	w2.Wait()
	demo3()
}

func demo3() {
	var c1 = make(chan int)
	var c2 = make(chan int)
	//var c1 = make(chan int, 1)
	//var c2 = make(chan int, 1)
	go demo3SUb(c1, c2)
	c1 <- 1
	fmt.Println(<-c2)
	close(c1)
	close(c2)

}

//只写 channel
func demo3SUb(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v + 1
	}
}
