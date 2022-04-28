package sorting

import "fmt"

func SelectionSort(arr []int) []int {
	var lowestValueIndex int
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		lowestValueIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lowestValueIndex] {
				lowestValueIndex = j
			}
			count++
		}
		if lowestValueIndex != i {
			arr[i], arr[lowestValueIndex] = arr[lowestValueIndex], arr[i]
		}
	}
	fmt.Println("Number of Steps taken: ", count)
	return arr
}
