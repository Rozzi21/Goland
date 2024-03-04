package main

import "fmt"

func main() {
	fmt.Println(sumAll(123, 121, 1351, 513))
	fmt.Println(sumAll(1, 3, 4, 5, 5))

	numbers := []int{10, 10, 10, 10}

	total := sumAll(numbers...)

	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
