package sorting

import "fmt"

func BubbleSort(arr []int) []int {

	i := len(arr) - 1
	sorted := false
	count := 0
	for !sorted {
		sorted = true
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = false
			}
			count++
		}
		i--
	}
	fmt.Println("Number of Steps taken: ", count)
	return arr
}
