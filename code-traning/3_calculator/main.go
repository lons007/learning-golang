package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// var c float64

func add(a, b float64) float64 {
	result := a + b
	return result
}

func subtract(a, b float64) float64 {
	result := a - b
	return result
}

func multiply(a, b float64) float64 {
	result := a * b
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Ошибка, на 0 делить нельзя!")
	}
	result := a / b

	return result, nil
}

func getNumber(prompt string, reader *bufio.Reader) float64 {
	var number float64
	for {
		fmt.Println(prompt)
		prompt, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		prompt = strings.TrimSpace(prompt)
		number, err = strconv.ParseFloat(prompt, 64)
		if err != nil {
			fmt.Println("Введено не число, а символ:", prompt)
			continue
		}
		break
	}
	return number
}

func main() {
	var a float64
	var b float64
	var c float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Добро пожаловать в калькулятор!")

	calc := false
	for calc == false {
		var operator string
		a = getNumber("Введите число A:", reader)

		for {
			var err error
			fmt.Println("Введите математическое действие +, -, *, /")
			operator, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			operator = strings.TrimSpace(operator)
			if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
				fmt.Println("Ошибка, оператор введен не правильно! Повторите попытку")
				continue
			}
			break
		}

		b = getNumber("введите число B:", reader)

		if operator == "+" {
			c = add(a, b)
		}

		if operator == "-" {
			c = subtract(a, b)
		}

		if operator == "*" {
			c = multiply(a, b)
		}

		if operator == "/" {
			var err error
			c, err = divide(a, b)
			if err != nil {
				fmt.Println(err)
				c = math.NaN()
			}
		}

		fmt.Println("Результат операции A и B:", c)
		fmt.Println("Хотите выполнить еще одну операцию? (yes/no)")
		next, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)

		}
		next = strings.TrimSpace(next)
		if next != "yes" && next != "y" {
			fmt.Println("Спасибо,что воспользовались нашим калькулятором")
			calc = true
		}
	}
}
