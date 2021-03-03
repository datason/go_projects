package main

import (
	"fmt"
	"time"
)

func longSQLQuery() chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- true
	}()
	return ch
}

func main() {
	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout happened")
	case <-time.After(time.Second):
		fmt.Println("timer.After timeout happened")
	case result := <-longSQLQuery():
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result:", result)
	}
}
