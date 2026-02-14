package main

import (
	"fmt"
	"learning-golang/src/github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.City = "Omaha"
	subscriber.Street = "123 Oak st"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("PostalCode:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.City = "Portland"
	employee.Street = "456 Elm St"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("PostalCode:", employee.PostalCode)

}