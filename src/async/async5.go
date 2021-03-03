package main

import "fmt"

func example(in chan int) {
	val := <-in
	fmt.Println("GO: get from chan", val)
	fmt.Println("GO: after read from chan")
}

func f(out chan<- int) {
	for i := 0; i <= 10; i++ {
		fmt.Println("before", i)
		out <- i
		fmt.Println("after", i)
	}
	close(out)
	fmt.Println("generator finish")
}

func main() {
	in := make(chan int, 0)
	go f(in)

	for i := range in {
		fmt.Println("\tget", i)
	}
}
