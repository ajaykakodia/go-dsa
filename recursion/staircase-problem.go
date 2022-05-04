package main

var countSteps int

func staircase(destinationSteps int) int {
	if destinationSteps <= 0 {
		return 0
	} else if destinationSteps == 1 {
		return 1
	} else if destinationSteps == 2 {
		return 2
	} else if destinationSteps == 3 {
		return 4
	}
	return staircase(destinationSteps-1) + staircase(destinationSteps-2) + staircase(destinationSteps-3)
}
func staircase2(destinationSteps int) int {
	countSteps++
	if destinationSteps < 0 {
		return 0
	} else if destinationSteps == 1 || destinationSteps == 0 {
		return 1
	}

	return staircase2(destinationSteps-1) + staircase2(destinationSteps-2) + staircase2(destinationSteps-3)
}

func staircase3(destinationSteps int, dic map[int]int) int {
	countSteps++
	if destinationSteps < 0 {
		return 0
	} else if destinationSteps == 1 || destinationSteps == 0 {
		return 1
	}
	var tls, sls, ls int
	if _, ok := dic[destinationSteps-1]; !ok {
		dic[destinationSteps-1] = staircase3(destinationSteps-1, dic)
	}
	tls = dic[destinationSteps-1]

	if _, ok := dic[destinationSteps-2]; !ok {
		dic[destinationSteps-2] = staircase3(destinationSteps-2, dic)
	}
	sls = dic[destinationSteps-2]

	if _, ok := dic[destinationSteps-3]; !ok {
		dic[destinationSteps-3] = staircase3(destinationSteps-3, dic)
	}
	ls = dic[destinationSteps-3]

	return tls + sls + ls
}
