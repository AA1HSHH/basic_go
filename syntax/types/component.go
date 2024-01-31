package main

import "fmt"

type Outer struct {
	Inner
}
type Inner struct {
}

func (o Outer) name() string {
	return "Outer"
}
func (i Inner) name() string {
	return "Inner"
}
func (i Inner) sayHello() {
	fmt.Println("hello" + i.name())
}

func Components() {
	var o Outer
	o.sayHello()
}

func main() {
	Components()
}
