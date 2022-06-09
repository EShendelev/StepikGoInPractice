package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func shuffle(nums []int) {
	// Функция `Seed()` инициализирует генератор случайных чисел
	// здесь мы используем константу `42`, чтобы программу
	// можно было проверить тестами.
	// В реальных задачах не используйте константы!
	// Используйте, например, время в наносекундах:
	// rand.Seed(time.Now().UnixNano())
	// rand.Seed(time.Now().UnixNano())
	// перетасуйте `nums` с помощью `rand.Shuffle()``
	rand.Seed(42)
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
