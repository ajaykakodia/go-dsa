package main

func greatestProductofThreeNumberInArray(arr []int) (pr int) {
	crPr := 0
	for i := 0; i < len(arr)-3; i++ {
		crPr = arr[i] * arr[i+1] * arr[i+2]
		if crPr > pr {
			pr = crPr
		}
	}
	return
}
