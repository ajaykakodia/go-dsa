package main

import "math"

func GreedyStockProblem(arr []int) bool {
	smallestNumber := arr[0]
	middleNumber := math.MaxInt

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestNumber {
			smallestNumber = arr[i]
		} else if arr[i] > smallestNumber && arr[i] <= middleNumber {
			middleNumber = arr[i]
		} else if arr[i] > middleNumber {
			return true
		}
	}
	return false
}
