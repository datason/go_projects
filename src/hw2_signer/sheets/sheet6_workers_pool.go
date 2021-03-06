package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	goroutinesNum = 5
	iterationsNum = 7
)

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", goroutinesNum-in),
		"th", in,
		"iter", j, strings.Repeat("*", j))
}

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{}
	defer wg.Done()

	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
	}
	<-quotaCh
}

func main() {
	quotaLimit := 1
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1) // add worker
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}
