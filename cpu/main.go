package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	var result string
	result = calculateResult(100_100)

	fmt.Println("elapsed:", time.Since(start))
	fmt.Println("length:", len(strings.Split(result, ",")))
}
