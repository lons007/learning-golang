package main

import "fmt"

type Task struct {
	ID        int
	Title     string
	Completed bool
	Category  string
}

type CategoryStats struct {
	Name      string
	Total     int
	Completed int
}

func (t *Task) Complete() {
	t.Completed = true
}

func (c *CategoryStats) Percent() float64 {
	if c.Total != 0 {
		return (float64(c.Completed) / float64(c.Total)) * 100
	} else {
		return 0
	}
}

type Stable interface {
	Percent() float64
}

func CalculateStats(tasks []Task) map[string]CategoryStats {
	m := make(map[string]CategoryStats)

	for _, task := range tasks {
		category := task.Category

		stats, exists := m[category] 
		if !exists {
			var newStats CategoryStats
			newStats.Name = category
			newStats.Total = 1
			if task.Completed {
				newStats.Completed = 1
			} else {
				newStats.Completed = 0
			}

			m[category] = newStats
		} else {
			var updatedStats CategoryStats
			updatedStats.Name = stats.Name
			updatedStats.Total = stats.Total + 1
			if task.Completed {
				updatedStats.Completed = stats.Completed +1
			} else {
				updatedStats.Completed = stats.Completed
			}
			m[category] = updatedStats
		}
	}
	return m
}

func PrintReport(stats map[string]CategoryStats) {
	for key, value := range stats {
		fmt.Printf("Категория: %s\n", key)
		fmt.Printf(" Всего задач: %d\n", value.Total)
		fmt.Printf(" Выполнено задач: %d\n", value.Completed)
		fmt.Printf(" Процент: %0.2f%%\n", value.Percent())
	}
}


func main() {
	t1 := Task{ID: 1, Title: "Купить хлеб", Completed: false, Category: "Личное"}
	t2 := Task{ID: 2, Title: "Создать отчет", Completed: false, Category: "Работа"}
	t3 := Task{ID: 3, Title: "Позвонить маме", Completed: true, Category: "Личное"}
	t4 := Task{ID: 4, Title: "Закончить проект", Completed: true, Category: "Работа"}
	t5 := Task{ID: 5, Title: "Ответить на письма", Completed: true, Category: "Работа"}

	tasks := []Task{t1, t2, t3, t4, t5}

	stats := CalculateStats(tasks)
	fmt.Println(stats)

	PrintReport(stats)

}
