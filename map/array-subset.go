package main

func arraySubset(arr1, arr2 []string) bool {

	largeArrMap := make(map[string]bool, len(arr1))
	var smallArr, largeArr []string
	if len(arr1) > len(arr2) {
		largeArr = arr1
		smallArr = arr2
	} else {
		largeArr = arr2
		smallArr = arr1
	}

	for _, c := range largeArr {
		largeArrMap[c] = true
	}

	for _, c := range smallArr {
		if !largeArrMap[c] {
			return false
		}
	}

	return true
}
