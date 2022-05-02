package main

func mergeSoretedArray(arr1, arr2 []int) []int {
	arr3 := []int{}
	p1 := 0
	p2 := 0

	for p1 != len(arr1) && p2 != len(arr2) {
		if arr1[p1] < arr2[p2] {
			arr3 = append(arr3, arr1[p1])
			p1 += 1
			continue
		}
		arr3 = append(arr3, arr2[p2])
		p2 += 1
	}

	if p1 != len(arr1) {
		arr3 = append(arr3, arr1[p1:]...)
	}

	if p2 != len(arr2) {
		arr3 = append(arr3, arr2[p2:]...)
	}

	return arr3
}
