package main

import "fmt"

func main() {
	fmt.Println("main from recursion")

	fact := factorial(5)
	fmt.Println("factorial: ", fact)

	fileTraversal("../go-dsa")
	as := arraySum([]int{23, 21, 24, 25, 10})
	fmt.Println("arraySum: ", as)
}
