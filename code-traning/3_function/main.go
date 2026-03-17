package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var c float64

func add(a, b float64) float64 {
	c = a + b
	return c
}

func subtract(a, b float64) float64 {
	c = a - b
	return  c
}

func multiply(a,b float64) float64 {
	c = a * b
	return c
}

func divide(a,b float64) float64 {
	c = a / b
	return c
}

func main() {
	var c float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите число A:")
	inputA, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputA = strings.TrimSpace(inputA)

	a, err := strconv.ParseFloat(inputA, 64)
	if err != nil {
		fmt.Println("Введено не число, значение числа A будет заменено на 0")
	}

	fmt.Println("Введите матиматическое действие +, -, *, /")
	operator, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	operator = strings.TrimSpace(operator)

	fmt.Println("Введите число B:")
	inputB, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputB = strings.TrimSpace(inputB)

	b, err := strconv.ParseFloat(inputB, 64)
	if err != nil {
		fmt.Println("Введено не число, значение числа B будет заменено на 0")
	}

	if operator == "+" {
		c = add(a,b)
	}

	if operator == "-" {
		c = subtract(a,b)
	}

	if operator == "*" {
		c = multiply(a,b)
	}

	if operator == "/" {
		c = divide(a,b)
	}
	
	fmt.Println("Результат операции A и B:", c)
	fmt.Println("Хотите выполнить еще одну операцию? (yes/no)")
}
