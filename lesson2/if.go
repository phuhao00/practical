package main

import "fmt"

type Animal interface {
	Eat(interface{})
	Name() string
}
type AnimalCategory int

const (
	AnimalCategory1 AnimalCategory = iota + 1
	AnimalCategory2
)

type Base struct {
	Nick string `json:"nick"`
	Age  int    `json:"age"`
}

func (b *Base) Eat(i interface{}) {
	fmt.Println(" not clear ")
}

func (b *Base) Name() string {
	return ""
}
