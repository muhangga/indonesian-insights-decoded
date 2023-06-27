package main

import (
	"fmt"
)

func SumNumber(a, b int) int {
	return a + b
}

func main() {
	num1 := 5
	num2 := 10

	result := SumNumber(num1, num2)
	fmt.Println("Result: ", result)
}
