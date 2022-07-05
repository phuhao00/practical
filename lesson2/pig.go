package main

import "fmt"

type Pig struct {
	Nick string `json:"nick"`
	Age  int    `json:"age"`
}

func (p *Pig) Eat(i interface{}) {
	fmt.Println("cookie")
}

func (p *Pig) Name() string {
	return "peiqi"
}
