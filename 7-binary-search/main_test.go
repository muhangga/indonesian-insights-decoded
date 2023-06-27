package main

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	target := 10
	idx := binarySearch(arr, target)
	if idx != 4 {
		t.Errorf("Expected 4, got %d", idx)
	}

	target = 15
	idx = binarySearch(arr, target)
	if idx != -1 {
		t.Errorf("Expected -1, got %d", idx)
	}
}
