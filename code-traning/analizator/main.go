package main

import (
	"fmt"
	"strings"
)

type Word struct {
	Text     string
	Length   int
	IsUnique bool
}

type Analyzer interface {
	TotalWords() int
	UniqueWords() int
	AverageWordLength() float64
}

type TextStats []Word

func (t TextStats) TotalWords() int {
	return len(t)
}

func (t TextStats) UniqueWords() int {
	uniqueWord := 0
	for _, word := range t {
		if word.IsUnique {
			uniqueWord++
		}
	}
	return uniqueWord
}

func (t TextStats) AverageWordLength() float64 {
	if len(t) != 0 {
		lenAllWord := 0
		for _, word := range t {
			lenAllWord += word.Length
		}
		return float64(lenAllWord) / float64(len(t))
	}
	return 0.0
}

func analyzeText(text string) TextStats {
	var wordRepit = make(map[string]int)
	var slice TextStats
	words := strings.Fields(text)
	for _, s := range words {
		s = strings.Trim(s, ".,!-_;:")
		s = strings.ToLower(s)
		wordRepit[s]++
	}

	for _, s := range words {
		s = strings.Trim(s, ".,!-_;:")
		s = strings.ToLower(s)
		var word Word
		word.Text = s
		word.Length = len(s)
		if wordRepit[s] == 1 {
			word.IsUnique = true
		}
		slice = append(slice, word)
	}

	// fmt.Println("Дебаг wordRepit:", wordRepit)
	// fmt.Println("Дебаг slice:", slice)
	return slice
}

func printReport (a Analyzer) {
	fmt.Println("Статистика текста:")
	fmt.Printf("Всего слов: %d\n", a.TotalWords())
	fmt.Printf("Уникальных слов: %d\n", a.UniqueWords())
	fmt.Printf("Средняя длина слова: %0.2f\n",a.AverageWordLength())
}

func main() {
	text := "Go - это отличный язык программирования. Go простой, Go выразительный, Go эффективный!"
	stats := analyzeText(text)
	printReport(stats)

}
