package sorting

import "fmt"

func InsertionSort(arr []int) []int {

	var tempData int
	count := 0
	for i := 1; i < len(arr); i++ {
		tempData = arr[i]
		position := i - 1
		for position >= 0 {
			count++
			if arr[position] > tempData {
				arr[position+1] = arr[position]
				position--
				continue
			}
			break
		}
		arr[position+1] = tempData
	}
	fmt.Println("Number of Steps taken: ", count)
	return arr
}
