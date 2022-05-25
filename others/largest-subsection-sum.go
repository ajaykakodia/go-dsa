package main

func LargestSubsectionSum(arr []int) int {
	largestSum := arr[0]
	var currentSum int

	for i := 0; i < len(arr); i++ {
		currentSum = currentSum + arr[i]
		if currentSum > largestSum {
			largestSum = currentSum
		}
		if currentSum <= 0 {
			currentSum = 0
		}
	}

	return largestSum
}
