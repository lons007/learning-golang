package main

import "fmt"

type Resetter interface {
	Reset()
}

type Timer struct {
	seconds int
}

type Scope struct {
	points int
}

func (t *Timer) Reset() {
	t.seconds = 0
}

func (s *Scope) Reset() {
	s.points = 0
}

func ReserAll(resetters []Resetter) {
	for _, resetter := range resetters {
		resetter.Reset()
	}
}

func main() {
	secondsVar := Timer{seconds: 120}
	pointsVar := Scope{points: 50}

	slice := []Resetter{&secondsVar, &pointsVar}

	fmt.Println("До сброса:")
	fmt.Printf("Таймер: %d\n", secondsVar.seconds)
	fmt.Printf("Счет: %d\n", pointsVar.points)

	ReserAll(slice)

	fmt.Println("После сброса:")
	fmt.Printf("Таймер: %d\n", secondsVar.seconds)
	fmt.Printf("Счет: %d\n", pointsVar.points)

}
