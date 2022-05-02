package main

import (
	"fmt"
)

func main() {
	fmt.Println("BigO main method")

	arr := []int{30, 50, 70, 90, 30, 50, 70}
	fmt.Println("Array Init")
	res := hundredSumArray(arr)
	if res {
		fmt.Println("Passed array is hundredSumArray")
	} else {
		fmt.Println("Passed array is not hundredSumArray")
	}

	res1 := mergeSoretedArray([]int{1, 5, 8, 10, 12}, []int{2, 7, 9, 11, 13, 15, 17})
	fmt.Println("Merged sorted array: ", res1)

	res1 = mergeSoretedArray([]int{1, 5, 8, 10, 12, 19, 21, 24, 25, 29}, []int{2, 7, 9, 11, 13, 15, 17})
	fmt.Println("Merged sorted array: ", res1)

	res1 = mergeSoretedArray([]int{}, []int{2, 7, 9, 11, 13, 15, 17})
	fmt.Println("Merged sorted array: ", res1)

	res = findNeedleInHayStack("abcdefghi", "def")
	if res {
		fmt.Println("needle found")
	} else {
		fmt.Println("needle not found")
	}

	res = findNeedleInHayStack("abcdefghi", "dd")
	if res {
		fmt.Println("needle found")
	} else {
		fmt.Println("needle not found")
	}

	res = findNeedleInHayStack("abcdefghi", "hiz")
	if res {
		fmt.Println("needle found")
	} else {
		fmt.Println("needle not found")
	}

	res = findNeedleInHayStack2("abcdefghi", "def")
	if res {
		fmt.Println("findNeedleInHayStack2 needle found")
	} else {
		fmt.Println("findNeedleInHayStack2 needle not found")
	}

	res = findNeedleInHayStack2("abcdefghi", "dd")
	if res {
		fmt.Println("findNeedleInHayStack2 needle found")
	} else {
		fmt.Println("findNeedleInHayStack2 needle not found")
	}

	res = findNeedleInHayStack2("abcdefghi", "hiz")
	if res {
		fmt.Println("findNeedleInHayStack2 needle found")
	} else {
		fmt.Println("findNeedleInHayStack2 needle not found")
	}

	res = findNeedleInHayStack2("abcdefghi", "abc")
	if res {
		fmt.Println("findNeedleInHayStack2 needle found")
	} else {
		fmt.Println("findNeedleInHayStack2 needle not found")
	}

	res2 := greatestProductofThreeNumberInArray([]int{2, 5, 8, 1, 2, 5, 7, 8, 11, 2})

	fmt.Println("greatestProductofThreeNumberInArray: ", res2)

	res3 := luckyResume([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println("luckyResume Number : ", res3)
}
