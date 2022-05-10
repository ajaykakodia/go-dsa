package main

import "fmt"

func main() {
	fmt.Println("main from recursion")

	fact := factorial(5)
	fmt.Println("factorial: ", fact)

	// fileTraversal("../go-dsa")

	as := arraySum([]int{23, 21, 24, 25, 10})
	fmt.Println("arraySum: ", as)

	sum := staircase(11)
	fmt.Println("staircase count: ", sum)

	sum = staircase2(11)
	fmt.Println("staircase2 count: ", sum)
	fmt.Println("staircase2 countSteps: ", countSteps)

	dic := make(map[int]int)
	countSteps = 0
	sum = staircase3(11, dic)
	fmt.Println("staircase3 count: ", sum)
	fmt.Println("staircase3 countSteps: ", countSteps)

	anagrams := anagramGen("abc")
	fmt.Println("anagramGen : ", anagrams)
	anagrams = anagramGen("abcd")
	fmt.Println("anagramGen : ", anagrams)

	charCount := countCharInArray([]string{"abc", "f", "fhgy", "dfdf", "fgdfghghgfhgfh"})
	fmt.Println("countCharInArray: ", charCount)

	trnum := triangularNumber(8)
	fmt.Println("triangularNumber : ", trnum)

	evenarrays := evenArray([]int{3, 4, 6, 8, 7, 9})
	fmt.Println("evenArray : ", evenarrays)

	fxIndex := findFirstX("abcfghdgexdfgdfgdgfgerty")
	fmt.Println("findFirstX : ", fxIndex)

	up := uniquePath(3, 7)
	fmt.Println("uniquePath : ", up)

	countSteps = 0
	fib := fibonacciSeries(21)

	fmt.Printf("fibonacciSeries Number : %d for sequence 21 and stepCounts : %d \n", fib, countSteps)

	countSteps = 0
	fib = fibonacciSeriesMemoization(21, make(map[int]int))

	fmt.Printf("fibonacciSeriesMemoization Number : %d for sequence 21 and stepCounts : %d \n", fib, countSteps)

	countSteps = 0
	fib = fibonacciSeriesBottomUp(21)

	fmt.Printf("fibonacciSeriesBottomUp Number : %d for sequence 21 and stepCounts : %d \n", fib, countSteps)
}
