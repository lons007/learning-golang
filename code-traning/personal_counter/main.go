package main

import "fmt"

func main() {
	var userName string
	var userAge int
	var userRate float64

	userName = "Ник"
	userAge = 18
	userRate = 5.51

	var newAge int
	newAge = userAge + 10

	fmt.Printf("Привет, %s\n", userName)
	fmt.Printf("Тебе %d лет\n", userAge)
	fmt.Printf("Твой рейтинг: %0.1f\n", userRate)
	fmt.Printf("Через 10 лет тебе будет %d лет.", newAge)

}