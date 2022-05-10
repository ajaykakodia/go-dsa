package main

import (
	"fmt"

	"github.com/ajaykakodia/go-dsa/searching"
	"github.com/ajaykakodia/go-dsa/sorting"
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
	// Bubble Sort
	arr = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19}
	sa := sorting.BubbleSort(arr)
	fmt.Println("Bubble Sorted Array: ", sa)

	// Selection Sort
	arr = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	sa = sorting.SelectionSort(arr)
	fmt.Println("Selection Sorted Array: ", sa)

	// Insertion Sort
	arr = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	sa = sorting.InsertionSort(arr)
	fmt.Println("InsertionSort Array: ", sa)

	// Quick Sort
	arr = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	qs := sorting.QuickSort{}
	qs.SetArray(arr)
	qs.SortArray()

	fmt.Println("QuickSort Array: ", qs.SortedArray())

	arr1 := []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	qs.SetArray(arr1)
	val := qs.QuickSelect(9)
	fmt.Println("QuickSelect Kth Value: ", val)

}
