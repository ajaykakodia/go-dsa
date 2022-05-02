package main

func luckyResume(arr []int) int {
	startSide := "top"
	for len(arr) != 1 {
		hl := len(arr) / 2
		if startSide == "top" {
			arr = arr[:hl]
			startSide = "bottom"
			continue
		}
		arr = arr[hl:]
		startSide = "top"
	}
	return arr[0]
}
