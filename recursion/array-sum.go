package main

func arraySum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + arraySum(arr[1:])
}
