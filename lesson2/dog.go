package main

import "fmt"

type Dog struct {
	Nick string `json:"nick"`
	Age  int    `json:"age"`
}

func (d *Dog) Eat(i interface{}) {
	fmt.Printf("nick is:%v \n", d.Nick)
}

func (d *Dog) Name() string {
	return d.Nick
}
