package main

import "fmt"

// hundredSumArray
func hundredSumArray(arr []int) bool {
	li := 0
	ri := len(arr) - 1
	fmt.Println("hundredSumArray called")
	for li < len(arr)/2 {
		if arr[li]+arr[ri] != 100 {
			return false
		}
		li += 1
		ri -= 1
	}
	return true
}
