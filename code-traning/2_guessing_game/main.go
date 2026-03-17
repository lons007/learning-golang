package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	guess := 7
	reader := bufio.NewReader(os.Stdin)

	gameOver := false

	fmt.Println("Игра угадай число")
	fmt.Println("При вводе 0 игра завершится")

	for gameOver == false {
		fmt.Println("Угадайте число от 1 до 10:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		inputGuess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка, введите число!")
			continue
		}

		if inputGuess == 0 {
			fmt.Println("Игра завершена")
			gameOver = true
			break
		} else if inputGuess < 1 || inputGuess > 10 {
			fmt.Println("Число должно быть в диапозоне от 1 до 10")
			continue
		} else if inputGuess > guess {
			fmt.Println("Слишком много!")
		} else if inputGuess < guess {
			fmt.Println("Слишком мало!")
		} else if inputGuess == guess {
			fmt.Println("Вы угадали!")
			fmt.Println("Если хотите сыграть еще раз введите: yes")

			answer, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			answer = strings.TrimSpace(answer)

			if answer != "yes" {
				fmt.Println("Спасибо за игру!")
				gameOver = true
			}
		}
	}
}
