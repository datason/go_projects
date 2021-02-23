package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iteratationsNum = 7
	goroutinesNum   = 5
)

func doSomeWork(in int) {
	for j := 0; j < iteratationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func main1() {
	for i := 0; i < goroutinesNum; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*", strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}

func main2() {
	ch1 := make(chan int, 2)

	go func(in chan int) {
		fmt.Println(in)
		val := <-in
		fmt.Println("GO: get from chan", val)
		fmt.Println("GO: after read from chan")
	}(ch1)

	ch1 <- 42
	ch1 <- 45
	ch1 <- 234
	fmt.Println("MAIN: after put to chan")
}

func a() {
	for i := 0; i < 5; i++ {
		fmt.Println("a")
	}
}

func b() {
	for i := 0; i < 5; i++ {
		fmt.Println("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}
