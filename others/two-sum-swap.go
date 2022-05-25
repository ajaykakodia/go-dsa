package main

func TwoSumSwap(arr1, arr2 []int) []int {
	swappedIndex := []int{}

	sum1, sum2 := 0, 0
	bigArrayCounterMap := make(map[int]int)

	for i, val := range arr1 {
		sum1 += val
		bigArrayCounterMap[val] = i
	}

	for _, val := range arr2 {
		sum2 += val
	}

	shiftAmount := (sum1 - sum2) / 2

	for i, val := range arr2 {
		if firstIndex, okk := bigArrayCounterMap[val+shiftAmount]; okk {
			swappedIndex = append(swappedIndex, firstIndex, i)
		}
	}

	return swappedIndex
}
