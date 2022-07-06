package main

import (
	"fmt"
	"sync"
)

var ch2 = make(chan int)    //无缓冲
var ch3 = make(chan int, 5) //有缓冲
var w sync.WaitGroup

func ChDemo() {
	w.Add(1)
	c := ChDemoSub()
	go func() {
		w.Wait()
		fmt.Println(<-c)
		close(c)
		w2.Done()
	}()
}

func ChDemoSub() chan int {
	var ch1 chan int
	ch1 = make(chan int, 2) //有缓冲
	ch1 <- 9
	w.Done()
	return ch1
}

func ChDemo2() {
	ch3 <- 7
	go ChDemo2SUb()
}

func ChDemo2SUb() {
	val := <-ch3
	fmt.Println(val)
	close(ch3)
	w2.Done()
}
