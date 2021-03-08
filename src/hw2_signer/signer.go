package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Message struct {
	str string
	id  int
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func invokerDataSignerMd5(output_channel chan string, input string, quotaCh chan struct{}) {
	quotaCh <- struct{}{}
	output_channel <- DataSignerMd5(input)
	<-quotaCh
}

func invokerDataSignerCrc32(output_channel chan string, input string) {
	output_channel <- DataSignerCrc32(input)
}

func singleHash(output_channel chan string, input string, quotaCh chan struct{}) {
	md5_channel := make(chan string)
	crc32_channel := make(chan string)
	crc32_channel_add := make(chan string)
	go invokerDataSignerMd5(md5_channel, input, quotaCh)
	go invokerDataSignerCrc32(crc32_channel_add, input)
	go invokerDataSignerCrc32(crc32_channel, <-md5_channel)
	output_channel <- <-crc32_channel_add + "~" + <-crc32_channel
}

func invokerDataSignerCrc32ForMessage(output_channel chan Message, input Message) {
	fmt.Printf("Run routine %d \n", input.id)
	msg := Message{DataSignerCrc32(input.str), input.id}
	output_channel <- msg
}

func multiHash(outputChannel chan string, input string) {
	thetas := []string{"0", "1", "2", "3", "4", "5"}
	channel := make(chan Message, len(thetas))
	// run goroutines
	for id, theta := range thetas {
		go invokerDataSignerCrc32ForMessage(channel, Message{theta + input, id})
	}
	// collect results
	var messages []Message
	for i := 0; i < len(thetas); i++ {
		messages = append(messages, <-channel)
	}
	// sorted by index
	sort.SliceStable(messages, func(i, j int) bool {
		return messages[i].id < messages[j].id
	})
	// concatenate
	var result string
	for i := 0; i < len(thetas); i++ {
		result = result + messages[i].str
	}
	// send by channel
	outputChannel <- result
}

func combineResults(results []string) string {
	sort.Strings(results)
	return strings.Join(results[:], "_")
}

// func main() {
// 	defer elapsed("Time spent:")()
// 	// quotaCh := make(chan struct{}, 1)
// 	dataChannel := make(chan string)

// 	multiHash(dataChannel, "rabbit")
// 	fmt.Println(<-dataChannel)
// }

// Переделать Execute Pipeline так, чтобы работали
// type job func(in, out chan interface{})
func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	out := make(chan interface{})

	for i := 0; i < len(jobs); i++ {
		go jobs[i](in, out)
		in := out
		out := make(chan interface{})
	}

}

// job(func(in, out chan interface{}) {
// 	out <- uint32(1)
// 	out <- uint32(3)
// 	out <- uint32(4)
// }),
// job(func(in, out chan interface{}) {
// 	for val := range in {
// 		out <- val.(uint32) * 3
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }),
// job(func(in, out chan interface{}) {
// 	for val := range in {
// 		fmt.Println("collected", val)
// 		atomic.AddUint32(&recieved, val.(uint32))
// 	}
// })

// func _main() {
// 	defer elapsed("Time spent:")()
// 	quotaCh := make(chan struct{}, 1)
// 	data_channel := make(chan string)

// 	strs := []string{"0", "1"}

// 	// data_channel <- strs[0]
// 	// fmt.Println(<-data_channel)

// 	for _, s := range strs {
// 		// fmt.Printf("%T", s)
// 		go singleHash(data_channel, s, quotaCh) //, quotaCh
// 		fmt.Println(s, <-data_channel)
// 	}

// }
