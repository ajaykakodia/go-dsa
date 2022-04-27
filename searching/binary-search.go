package searching

import "fmt"

func BinarySearch(arr []int, key int) int {
	leftIndex := 0
	rightIndex := len(arr) - 1
	midIndex := (leftIndex + rightIndex) / 2

	for leftIndex < rightIndex {

		if arr[midIndex] == key {
			return midIndex
		} else if key < arr[midIndex] {
			rightIndex = midIndex - 1
		} else {
			leftIndex = midIndex + 1
		}
		midIndex = (leftIndex + rightIndex) / 2
		fmt.Println("step count...")
	}
	return -1
}
