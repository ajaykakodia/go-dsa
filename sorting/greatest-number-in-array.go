package sorting

import "sort"

func GreatestNumberInArray1(arr []int) int {

	gtNum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > gtNum {
			gtNum = arr[i]
		}
	}

	return gtNum
}

func GreatestNumberInArray2(arr []int) int {

	sort.Ints(arr)

	return arr[len(arr)-1]
}

func GreatestNumberInArray3(arr []int) int {

	isGreatestNumber := false
	for i := 0; i < len(arr); i++ {
		isGreatestNumber = true
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				isGreatestNumber = false
			}
		}

		if isGreatestNumber {
			return arr[i]
		}
	}

	return -1
}
