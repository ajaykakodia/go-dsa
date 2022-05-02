package main

func intersection(arr1, arr2 []int) []int {
	arr3 := make([]int, 0)
	am1 := make(map[int]bool, len(arr1))

	for _, val := range arr1 {
		am1[val] = true
	}

	for _, val := range arr2 {
		if am1[val] {
			arr3 = append(arr3, val)
		}
	}

	return arr3
}
