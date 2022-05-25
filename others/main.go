package main

import "fmt"

func main() {
	fmt.Println("Message from others....")

	reverseArray := ReverseArray([]int{65, 45, 23, 45, 78, 90, 12})
	fmt.Println("Reverse Array :", reverseArray)

	exists := TwoSum([]int{2, 0, 4, 1, 7, 9}, 10)
	fmt.Println("Two sum exists :", exists)

	gameWinner := GameWinner(12, "you")
	fmt.Println("Game Winner :", gameWinner)

	gameWinner = GameWinner2(10)
	fmt.Println("Game Winner 2 :", gameWinner)

	reverseArray = TwoSumSwap([]int{5, 3, 2, 9, 1}, []int{1, 12, 5})
	fmt.Println("TwoSumSwap switch required for Indies :", reverseArray)

	sum := LargestSubsectionSum([]int{3, -4, 4, -3, 5, -9})
	fmt.Println("LargestSubsectionSum :", sum)

	upwardTrendExists := GreedyStockProblem([]int{22, 25, 21, 18, 19, 17, 16, 20})
	fmt.Println("GreedyStockProblem upward trend exists:", upwardTrendExists)

	anagramCheck := AnagramChecker("startle", "rattles")
	fmt.Println("AnagramChecker :", anagramCheck)
}
