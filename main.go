package main

import (
	"fmt"

	"github.com/ajaykakodia/go-dsa/searching"
)

func main() {
	fmt.Println("first method to call")
	arr := []int{2, 4, 7, 34, 56, 67, 78, 89, 90, 100, 101, 123}
	index := searching.BinarySearch(arr, 7)
	if index == -1 {
		fmt.Println("Number is not exists in current array ")
		return
	}
	fmt.Println("found at: ", index)
}
