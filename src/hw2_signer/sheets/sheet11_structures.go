package main

import "fmt"

func double_int(i *int) {
	fmt.Println(i, *i)
	*i *= 2
	fmt.Println(i, *i)
}

type subscriber struct {
	name string
	rate float64
}

func applyDiscont(s *subscriber) {
	s.rate = 4
	fmt.Println(s.rate)
}

func main() {
	var s subscriber
	s.name = "John Doe"
	s.rate = 64
	fmt.Println(s.rate)
	applyDiscont(&s)
	fmt.Println(s.rate)
}

// func main() {
// 	i := 6
// 	double_int(&i)
// 	fmt.Println(i)
// }
