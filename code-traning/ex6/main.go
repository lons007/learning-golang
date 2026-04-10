package main

import "fmt"

type Sale struct {
	Product  string
	Amount   int
	Category string
}

type CategoryReport struct {
	Name       string
	TotalSales int
	Count      int
	Average    float64
}

func (c *CategoryReport) CalculateAverage() {
	c.Average = float64(c.TotalSales) / float64(c.Count)
}

func AnalyzeSales(sales []Sale) map[string]CategoryReport {

	m := make(map[string]CategoryReport)

	for _, sale := range sales {
		key, exist := m[sale.Category]
		
		// fmt.Printf("RUN key: %v exist: %v, map: %v \n", key, exist, m)
		fmt.Println(m)
		if !exist {
			key.Name = sale.Category
		}
		key.Count += 1
		key.TotalSales += sale.Amount
		key.CalculateAverage()
		m[sale.Category] = key

		
		
		// fmt.Printf("END key: %v exist: %v map: %v \n", key, exist, m )
		// fmt.Println("DEBUG", m)
	}
	return m
}

func PrintReport(report map[string]CategoryReport) {
	for key, value := range report {
		fmt.Printf("Категория: %s\n", key)
		fmt.Printf(" Продаж: %d\n", value.Count)
		fmt.Printf(" Общая сумма: %d\n", value.TotalSales)
		fmt.Printf(" Средний чек: %0.2f\n", value.Average)
	}
}

func main() {
	t1 := Sale{Product: "Ноутбук", Amount: 50000, Category: "Электроника"}
	t2 := Sale{Product: "Мышь", Amount: 1500, Category: "Электроника"}
	t6 := Sale{Product: "Книга Python", Amount: 700, Category: "Книги"}
	t3 := Sale{Product: "Клавиатура", Amount: 3000, Category: "Электроника"}
	t4 := Sale{Product: "Книга Go", Amount: 800, Category: "Книги"}
	t5 := Sale{Product: "Книга Python", Amount: 700, Category: "Книги"}

	sales := []Sale{t1, t2, t3, t4, t5,t6}

	report := AnalyzeSales(sales)
	PrintReport(report)
}
