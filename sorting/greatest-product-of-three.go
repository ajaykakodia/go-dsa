package sorting

import "sort"

var StepsTaken int

func GreatestProductOfThree(arr []int) int {
	gprod := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j {
				for k := 0; k < len(arr); k++ {
					if k != i && k != j {
						prod := arr[i] * arr[j] * arr[k]
						if prod > gprod {
							gprod = prod
						}
						StepsTaken++
					}
				}
			}
		}
	}

	return gprod
}

func GreatestProductOfThree2(arr []int) int {
	gprod := 0

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if i != j {
				for k := j + 1; k < len(arr); k++ {
					if k != i && k != j {
						prod := arr[i] * arr[j] * arr[k]
						if prod > gprod {
							gprod = prod
						}
						StepsTaken++
					}
				}
			}
		}
	}

	return gprod
}

func GreatestProductOfThree3(arr []int) int {
	gprod := 0

	sort.Ints(arr)

	l := len(arr) - 1
	gprod = arr[l-2] * arr[l-1] * arr[l]

	return gprod
}
