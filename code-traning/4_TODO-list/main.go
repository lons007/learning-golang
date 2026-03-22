package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func inputCommand() {
	fmt.Printf("Введите команду: add, edit, del, done, list, stats, exit\n")
}

func add(prompt string, reader *bufio.Reader, tasks []Task) Task {
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

func idTasks(prompt string, reader *bufio.Reader) (int, error) {
	for {

		fmt.Printf("%s ", prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("Ошибка чтения ввода: %w", err)
		}
		input = strings.TrimSpace(input)

		idTask, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом. Попробуйте снова.\n", input)
			continue
		}
		return idTask, nil
	}
}

func main() {
	var tasks []Task
	fmt.Println("Добро пожаловать в Трекер задач!")
	reader := bufio.NewReader(os.Stdin)
	commandExit := false
mainForlabel:
	for commandExit == false {
		inputCommand()
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		switch command {
		case "add":
			addTask := add("Введите название задачи:", reader, tasks)
			tasks = append(tasks, addTask)
			fmt.Printf("Задача %s добавлена с ID %d\n", addTask.Title, addTask.ID)
		case "done":
			id, err := idTasks("Введите id выполненной команды:", reader)
			if err != nil {
				log.Fatal(err)
			}
			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					fmt.Printf("Задача '%s' отмечена как выполненная\n", tasks[i].Title)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Задача с ID '%d' не найдена\n", id)
			}
		case "list":
			lenSlice := len(tasks)
			if lenSlice == 0 {
				fmt.Println("Список задач пуст")
			}
			for _, task := range tasks {
				var completed string
				if task.Completed {
					completed = "[✓]"
				} else {
					completed = "[ ]"
				}
				fmt.Printf("%d. %s %s\n", task.ID, completed, task.Title)
			}

		case "stats":
			allTasks := len(tasks)
			fmt.Printf("Всего задач: %d\n", allTasks)
			completed := 0
			for _, task := range tasks {
				if task.Completed {
					completed++
				}
			}
			fmt.Printf("Выполнено задач: %d\n", completed)
			notCompleted := allTasks - completed
			fmt.Printf("Не выполнено задач: %d\n", notCompleted)
			var progress float64 = 0
			if allTasks != 0 {
				progress = (float64(completed) / float64(allTasks)) * 100
			}
			fmt.Printf("Прогресс %0.2f%%\n", progress)

		case "del":
			delId, err := idTasks("Введите номер ID которую необходимо удалить:", reader)
			if err != nil {
				log.Fatal(err)
			}
			var delTitle string
			var result []Task
			taskFound := false
			for _, task := range tasks {
				if task.ID != delId {
					result = append(result, task)
				} else {
					delTitle = task.Title
					taskFound = true
				}
			}
			if taskFound {
				tasks = result
				fmt.Printf("Задача %s удалена\n", delTitle)
			} else {
				fmt.Printf("Задача %d не найдена\n", delId)
			}

		case "edit":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
				continue mainForlabel
			}
			id, err := idTasks("Введи ID задачи для редактирования: ", reader)
			if err != nil {
				log.Fatal(err)
			}
			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					oldTask := tasks[i].Title
					for {
						fmt.Println("Введи новое название задачи:")
						input, err := reader.ReadString('\n')
						if err != nil {
							log.Fatal(err)
						}
						input = strings.TrimSpace(input)

						if input == "" {
							fmt.Println("Значение не должно быть пустым! Повторите ввод")
							continue
						}

						tasks[i].Title = input
						fmt.Printf("Задача '%s' изменена на '%s'\n", oldTask, tasks[i].Title)
						found = true
						break
					}
				}

			}
			if !found {
				fmt.Println("Задача не найдена")
			}

		case "exit":
			fmt.Println("До встречи!")
			commandExit = true
		default:
		}

	}

}
