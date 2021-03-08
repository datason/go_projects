package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	channel := make(chan int)
	go func() {
		for _, n := range nums {
			channel <- n
		}
		close(channel) // to close channel in for cycle of reading from it
	}()
	return channel
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read in
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged ouput from c1 and c2.
	for n := range merge(c1, c2) {

	}
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Copy values from input channel to out
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	// Run goroutines to redirect all input channels to one general
	wg.Add(len(cs))
	for c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func main_1() {
	c := gen(2, 3)

	for n := range sq(sq(c)) {
		fmt.Println(n)
	}
}
