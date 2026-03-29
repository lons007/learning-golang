package main

import "fmt"

func groupByLength(words []string) map[int][]string {
	wordMap := make(map[int][]string)
	for _, word := range words {
		key := len(word)
		if value, ok := wordMap[key]; !ok {
			var slice []string
			slice = append(slice, word)
			wordMap[key] = slice
		} else {
			value = append(value, word)
			wordMap[key] = value
		}

	}
	return wordMap
}

func main() {
	slice := []string{"cat", "dog", "elephant", "bat", "rat", "tiger", "lion"}
	wordMap := groupByLength(slice)
	for key, value := range wordMap {
		fmt.Printf("Длина %d: %v\n", key, value)
	}
}
