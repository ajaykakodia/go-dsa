package main

// Provide highest product of any two number and can number can be negative

func HighestProductInArray(arr []int) int {

	var maxHighestNum, secondMax, maxMin, secondMin int

	for i := 0; i < len(arr); i++ {
		if arr[i] > maxHighestNum {
			secondMax = maxHighestNum
			maxHighestNum = arr[i]
			continue
		}
		if arr[i] > secondMax {
			secondMax = arr[i]
			continue
		}

		if arr[i] < maxMin {
			secondMin = maxMin
			maxMin = arr[i]
			continue
		}

		if arr[i] < secondMin {
			secondMin = arr[i]
		}
	}

	maxMinprod := secondMin * maxMin
	maxProd := secondMax * maxHighestNum

	if maxMinprod > maxProd {
		return maxMinprod
	}

	return maxProd
}
