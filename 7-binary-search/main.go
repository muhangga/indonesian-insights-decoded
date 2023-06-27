package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else if arr[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func isTargetExist(arr []int, target int) {
	idx := binarySearch(arr, target)
	if idx != -1 {
		fmt.Printf("Target found at index %d\n", idx)
	} else {
		fmt.Printf("Target not found, %d\n", idx)
	}
}
func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	target := 10
	isTargetExist(arr, target)

	target = 15
	isTargetExist(arr, target)
}
