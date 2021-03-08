package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	landmark := geo.Landmark{}

	err := landmark.SetName("North pole")
	if err != nil {
		log.Fatal(err)
	}

	err = landmark.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}

	err = landmark.SetLongitude(137.42)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(landmark.Latitude())
	fmt.Println(landmark.Longitude())
	fmt.Println(landmark.Name())
}
