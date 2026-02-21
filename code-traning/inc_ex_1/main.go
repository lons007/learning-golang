package main

import (
	"fmt"
	"learning-golang/src/github.com/headfirstgo/geo"
	"log"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}
