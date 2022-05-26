package main

func LongestConsecutiveSequence(arr []int) []int {
	longestSequence := []int{}

	arrMap := make(map[int]bool)

	for _, val := range arr {
		arrMap[val] = true
	}

	for _, val := range arr {
		currentSequence := []int{}
		_, ok := arrMap[val]
		for ok {
			currentSequence = append(currentSequence, val)
			val = val + 1
			_, ok = arrMap[val]
		}

		if len(currentSequence) > len(longestSequence) {
			longestSequence = currentSequence
		}
	}

	return longestSequence
}
