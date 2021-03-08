package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(1230)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())

	// DATE OF EVENT

	event := calendar.Event{}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(272)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("Mom's day")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())

}
