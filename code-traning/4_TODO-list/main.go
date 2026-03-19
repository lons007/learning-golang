package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func add(prompt string, reader *bufio.Reader, tasks[] Task) Task {
	var addTask Task
	fmt.Println(prompt)
	prompt, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	prompt = strings.TrimSpace(prompt)
	addTask.Title = prompt
	addTask.ID = len(tasks) + 1
	addTask.Completed = false
	return addTask
}

func main() {
	var tasks []Task
	fmt.Println("Добро пожаловать в Трекер задач!")
	fmt.Println("Доступные команды: add, done, list, stats, exit")
	reader := bufio.NewReader(os.Stdin)
	commandExit := false
	for commandExit == false {
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		switch command {
		case "add":
			addTask := add("Введите название задачи:", reader, tasks)
			tasks = append(tasks, addTask)
			fmt.Printf("Задача %s добавлна с ID %d\n", addTask.Title, addTask.ID)
			fmt.Println("DEBUG slice: ",tasks)

		case "done":
			fmt.Printf("Введена команда: %s\n", command)
		case "list":
			fmt.Printf("Введена команда: %s\n", command)
		case "stats":
			fmt.Printf("Введена команда: %s\n", command)
		case "exit":
			fmt.Println("До встречи!")
			commandExit = true
		default:
			fmt.Printf("Введена команду: add, done, list, stats, exit\n")
		}

	}

}
