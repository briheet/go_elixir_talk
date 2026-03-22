package main

import "fmt"

func main() {

	var answer int64

	numbers := []int64{1, 2, 3, 4}
	answer = sum(numbers)

	fmt.Println(answer)
}

func sum(numbers []int64) int64 {

	var sum int64
	for _, num := range numbers {
		sum += num
	}

	return sum
}
