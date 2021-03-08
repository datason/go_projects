package main

import (
	"calc"
	"fmt"
	"salut"
)

func main() {
	salut.Hello()
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calc.Subtract(7, 3))
	salut.Hi()
}
