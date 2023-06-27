package main

import "testing"

func TestSum(t *testing.T) {
	num1 := 5
	num2 := 10

	result := SumNumber(num1, num2)
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
}
