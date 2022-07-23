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

	done := make(chan struct{})
	counted := make(chan pair)

	// начало решения

	// запустите четыре горутины countWords()
	// вместо одной
	for i := 0; i < 4; i++ {
		go countWords(done, pending, counted)
		//fmt.Println(i+1, "Горутина")
	}
	// используйте канал завершения, чтобы дождаться
	// окончания обработки и закрыть канал counted
	go func() {
		for i := 0; i < 4; i++ {
			<-done
			//fmt.Println(i+1, "done")
		}
		close(counted)
	}()
	// конец решения

	return fillStats(counted)
}

// submit	Words отправляет слова на подсчет
func submitWords(next nextFunc, out chan<- string) {
	for {
		word := next()
		if word == "" {
			break
		}
		out <- word
	}
	close(out)
}

// countWords считает цифры в словах
func countWords(done chan<- struct{}, in <-chan string, out chan<- pair) {
	// реализуйте логику подсчета цифр
	// с использованием каналов done, in и out
}

// fillStats готовит итоговую статистику
func fillStats(in <-chan pair) counter {
	stats := counter{}
	for p := range in {
		stats[p.word] = p.count
	}
	return stats
}

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
	phrase := "1 22 333 4444 55555 666666 7777777 88888888"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
