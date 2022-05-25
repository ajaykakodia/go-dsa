package main

func TwoSum(arr []int, sum int) bool {
	counterNumberMap := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if counterNumberMap[sum-arr[i]] {
			return true
		}
		counterNumberMap[i] = true
	}
	return false
}
