package main

import "fmt"

func groupByFirstLetter(words []string) map[rune][]string {
	m := make(map[rune][]string)
	for _, word := range words {
		firstLetter:= rune(word[0])
		// value, exist := m[rune(word[0])]
		// if !exist {
		// 	var str []string
		// 	str = append(str, word)
		// 	m[rune(word[0])] = str
		// } else {
		// 	value = append(value, word)
		// 	m[rune(word[0])] = value
		// }
		m[firstLetter] = append(m[firstLetter], word)
	}
	return m
}

func main() {
	words := []string{"apple", "banana", "apricot", "blueberry", "cherry"}
	w := groupByFirstLetter(words)
	fmt.Println("вход:", words)
	fmt.Println("выход:", w)
}
