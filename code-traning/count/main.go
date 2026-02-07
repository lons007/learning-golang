package main

import "fmt"

func main() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]string)
	for _, item := range data{
		fmt.Println(item)
		fmt.Println(counts)
		counts[item] = counts[item] + "test"
		// counts[item] = counts[item] + 1
		fmt.Println(counts)
	}
	fmt.Println(counts)

	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		fmt.Println(count, ok)
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %s\n", letter, count)
		}
	}
}