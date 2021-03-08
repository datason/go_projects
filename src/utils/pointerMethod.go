package main

import "fmt"

type MyType string

func (m MyType) method() {
	fmt.Println("Method with value receicer")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer to value receicer")
}

func main() {
	value := MyType("jo")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
}
