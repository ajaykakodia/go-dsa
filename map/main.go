package main

import "fmt"

func main() {
	fmt.Println("Main from map folder")

	res := arraySubset([]string{"a", "b", "d", "l", "f", "g"}, []string{"l", "f"})
	if res {
		fmt.Println("arraySubset - one array is subset to other array")
	} else {
		fmt.Println("arraySubset - both are different array!!!")
	}

	res = arraySubset([]string{"a", "b", "d", "l", "f", "g"}, []string{"l", "f", "z"})
	if res {
		fmt.Println("arraySubset - one array is subset to other array")
	} else {
		fmt.Println("arraySubset - both are different array!!!")
	}

	res1 := intersection([]int{5, 10, 15, 20, 25, 17, 19}, []int{20, 5, 17, 40, 23, 45})
	fmt.Println("intersection - intersection of both array", res1)

	res1 = intersection([]int{5, 10, 15, 20, 25, 17, 19}, []int{200, 50, 170, 400, 230, 450})
	fmt.Println("intersection -No intersection data", res1)

	res2 := firstDuplicateString([]string{"a", "b", "c", "d", "c", "e"})
	fmt.Println("firstDuplicateString - ", res2)

	res3 := missingAlphabet("the quick brown box jumps over a lazy dog")
	fmt.Printf("missingAlphabet - %s\n", string(res3))

	res4 := firstNonDuplicateCharacterInString("minimum")
	fmt.Printf("firstNonDuplicateCharacterInString - %s\n", res4)
}
