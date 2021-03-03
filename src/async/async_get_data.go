package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, reture comments")
		out <- "32 comments"
	}(result)
	return result
}

func getPage() {
	// heavy or long
	resultCh := getComments()

	// easy
	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	// Okay, now use results from heavy deals
	commentsData := <-resultCh
	fmt.Println("main goroutine:", commentsData)
}

func main() {
	getPage()
}
