package main

import (
	"fmt"
	"time"
)

func repeat(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main1() {
	go repeat("x")
	go repeat("y")
	time.Sleep(time.Second)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main2() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)

	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}

func reportNap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "sleeping")
		time.Sleep(1 * time.Second)
	}
	fmt.Println(name, "wakes up!")
}

func send(myChannel chan string) {
	reportNap("sending goroutine", 2)
	fmt.Println("***sending value***")
	myChannel <- "a"
	fmt.Println("***sending value***")
	myChannel <- "b"
}

func main() {
	myChannel := make(chan string)
	go send(myChannel)
	reportNap("receiving goroutine", 8)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
