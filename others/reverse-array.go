package main

func ReverseArray(arr []int) []int {
	l := len(arr) - 1
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i] = arr[l-i], arr[i]
	}

	return arr
}
