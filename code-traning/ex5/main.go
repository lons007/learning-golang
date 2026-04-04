package main

import "fmt"

func gcountOccurrences(items []string) map[string]int {
	m := make(map[string]int)
	for _, item := range items {
	// 	value, exist := m[item]
	// 	if !exist {
	// 		m[item] = 1
	// 	} else {
	// 		m[item] = value + 1
	// 	}
	m[item]++
	}
	return m
}

func main() {
	items := []string{"cat", "dog", "cat", "bird", "dog", "cat"}
	i := gcountOccurrences(items)
	fmt.Println("вход:", items)
	fmt.Println("выход:", i)
}
