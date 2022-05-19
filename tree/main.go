package main

import "fmt"

func main() {
	fmt.Println("Message from Tree main")

	tree := BinaryTree{}
	values := []int{1, 5, 9, 2, 4, 10, 6, 3, 8}
	for _, val := range values {
		tree.Insert(val)
	}
	tree.InOrderTraversal()
	fmt.Println()
	fmt.Println("----------------")
	tree.PreOrderTraversal()
	fmt.Println()
	fmt.Println("----------------")
	tree.PostOrderTraversal()
	node := tree.Search(5)
	fmt.Println()
	if node != nil {
		fmt.Println("Node value :", node.data, "lNode :", node.lNode)
	}
	tree.Delete(4)
	fmt.Println()
	fmt.Println("----------------")
	tree.PreOrderTraversal()
	fmt.Println()
	gtNum := tree.GreatestNumber()
	fmt.Println("Greatest number: ", gtNum)

}
