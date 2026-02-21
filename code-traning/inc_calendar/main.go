package main

import (
	"fmt"
	"learning-golang/src/github.com/headfirstgo/calendar"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2018)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
	// date = calendar.Date{year: 0, month: 0, day: -2}
	// fmt.Println(date)
}
