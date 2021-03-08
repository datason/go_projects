package main

import "fmt"

type Number int

func (n Number) add(m int) {
	fmt.Println(int(n), " + ", m, " = ", int(n)+m)
}

func main() {
	a := Number(123)
	a.add(10)
}
