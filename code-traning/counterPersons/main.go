package main

import "fmt"

type VisitorCounter struct {
	count int
	lastVisitor string
}

func NewVisitorCounter() *VisitorCounter {
	return &VisitorCounter{
		count: 0,
		lastVisitor: "",
	}
}

func (v *VisitorCounter) AddVisit(name string) {
	v.count++
	v.lastVisitor = name
}

func (v VisitorCounter) GetCount() int {
	return v.count
}

func (v VisitorCounter) GetLastVisitor() string {
	return v.lastVisitor
}

func main() {
	counter := NewVisitorCounter()

	persons := []string{"Анна", "Петр", "Мария"}
	for _, person := range persons {
		counter.AddVisit(person)
	}
	fmt.Println("Посещений:",counter.GetCount())
	fmt.Println("Последний постетитель:", counter.GetLastVisitor())
}