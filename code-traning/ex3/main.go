package main

import (
	"fmt"
	"strconv"
)

func countGrades(grades []int) map[string]int {
	m := make(map[string]int)
	for _, grade := range grades {
		g := strconv.Itoa(grade)
		// state, exist := m[g]
		// if !exist {
		// 	m[g] = 1
		// } else {
		// 	m[g] = state + 1
		// }
		m[g] = m[g] + 1

	}
	return m
}

func main() {
	grades := []int{5, 4, 3, 5, 5, 4, 2}
	count := countGrades(grades)
	fmt.Println("вход:",grades)
	fmt.Println("выход:", count)

}
