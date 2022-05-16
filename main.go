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

	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestProductOfThree(arr1)

	fmt.Printf("GreatestProductOfThree : %d and Steps Taken :%d\n", val, sorting.StepsTaken)

	sorting.StepsTaken = 0
	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestProductOfThree2(arr1)

	fmt.Printf("GreatestProductOfThree2 : %d and Steps Taken :%d\n", val, sorting.StepsTaken)

	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestProductOfThree3(arr1)

	fmt.Println("GreatestProductOfThree3 : ", val)

	arr1 = []int{9, 3, 2, 5, 6, 1, 0, 4, 7, 8, 12, 15, 10, 14, 13}
	val = sorting.MissingNumber(arr1)

	fmt.Println("MissingNumber : ", val)

	arr1 = []int{9, 3, 2, 5, 6, 1, 0, 4, 7, 8, 12, 15, 10, 14, 13}
	val = sorting.MissingNumber2(arr1)

	fmt.Println("MissingNumber2 : ", val)

	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestNumberInArray1(arr1)

	fmt.Println("GreatestNumberInArray1 : ", val)

	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestNumberInArray2(arr1)

	fmt.Println("GreatestNumberInArray2 : ", val)

	arr1 = []int{5, 2, 7, 9, 30, 1, 17, 20, 14, 4, 3, 19, 8}
	val = sorting.GreatestNumberInArray3(arr1)

	fmt.Println("GreatestNumberInArray3 : ", val)

}
