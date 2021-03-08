package main

import (
	"fmt"
	"geo"
)

func main() {
	location := geo.Coordinages{Latitude: 37.42, Longitude: 120}
	fmt.Println("lat", location.Latitude)
	fmt.Println("lon", location.Longitude)
}
