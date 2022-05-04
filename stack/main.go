package main

import "fmt"

func main() {
	fmt.Println("main from stack folder.")

	stack := Stack{}
	stack.Push(2)
	stack.Push(5)
	stack.Push(7)
	stack.Push(10)
	stack.Push(12)
	fmt.Println(stack.Size())
	fmt.Println(stack.arr)
	fmt.Println(stack.Seek(), stack.arr)
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Size())
	stack.Push(20)
	stack.Push(42)
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Pop(), stack.arr)
	fmt.Println(stack.Size())

	err := Linter("{()[]}")
	if err != nil {
		fmt.Println("linting error: ", err.Error())
	} else {
		fmt.Println("No linting error")
	}

	err = Linter("{()[]}}")
	if err != nil {
		fmt.Println("linting error: ", err.Error())
	} else {
		fmt.Println("No linting error")
	}

	err = Linter("{{()[]}]")
	if err != nil {
		fmt.Println("linting error: ", err.Error())
	} else {
		fmt.Println("No linting error")
	}

	rvStr := reverseString("abcde")
	fmt.Println("reverseString: ", rvStr)

	rvStr = reverseString2("abcdefghjklm")
	fmt.Println("reverseString2: ", rvStr)
}
