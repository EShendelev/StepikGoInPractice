package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc returns the next word from the generator
type nextFunc func() string

// counter stores the number of digits in each word.
// each key is a word and value is the number of digits in the word.
type counter map[string]int

// pair stores a word and the number of digits in it
type pair struct {
	word  string
	count int
}

// countDigitsInWords counts digits in words,
// fetching each word with the next() function
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

// начало решения

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, out chan string) {
	for word := next(); len(word) > 0; {
		out <- word
		word = next()
	}
	close(out)

}

// countWords считает цифры в словах
func countWords(in chan string, out chan pair) {
	for word := range in {
		count := countDigits(word)
		out <- pair{word, count}
	}
	close(out)
}

// fillStats готовит итоговую статистику
func fillStats(in chan pair) counter {
	// p := pair{}
	stats := counter{}
	for p := range in {
		stats[p.word] = p.count
	}
	return stats

}

// конец решения

// countDigits returns the number of digits in a string
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats prints words and their digit counts
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator returns a generator,
// which emits words from a phrase.
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
