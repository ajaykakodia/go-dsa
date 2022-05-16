package sorting

import "sort"

func MissingNumber(arr []int) int {

	sort.Ints(arr)

	i := 0
	for ; i < len(arr); i++ {
		if arr[i] != i {
			return i
		}
	}

	return -1
}

func MissingNumber2(arr []int) int {

	numslen := len(arr) + 1
	numMap := make(map[int]bool)
	for _, val := range arr {
		numMap[val] = true

	}

	for i := 0; i < numslen; i++ {
		if numMap[i] == false {
			return i
		}
	}

	return -1
}
