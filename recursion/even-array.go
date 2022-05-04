package main

func evenArray(arr []int) []int {
	earr := []int{}
	if len(arr) == 0 {
		return earr
	}
	if arr[0]%2 == 0 {
		earr = append(earr, arr[0])
		return append(earr, evenArray(arr[1:])...)
	}
	return append(earr, evenArray(arr[1:])...)
}
