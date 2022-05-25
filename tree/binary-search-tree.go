package main

import "fmt"

type BinaryTree struct {
	root *node
}

type node struct {
	lNode *node
	rNode *node
	data  int
}

func (bt *BinaryTree) Insert(nodeValue int) {
	if bt.root == nil {
		bt.root = &node{data: nodeValue}
		return
	}
	insert(bt.root, nodeValue)
}

func (bt *BinaryTree) Search(val int) *node {
	return search(bt.root, val)
}

func (bt *BinaryTree) InOrderTraversal() {
	inOrderTraversal(bt.root)
}

func (bt *BinaryTree) PreOrderTraversal() {
	preOrderTraversal(bt.root)
}

func (bt *BinaryTree) PostOrderTraversal() {
	postOrderTraversal(bt.root)
}

func (bt *BinaryTree) Delete(nodeVale int) {
	deleteNode(bt.root, nodeVale)
}

func (bt *BinaryTree) GreatestNumber() int {
	return greatestNumber(bt.root)
}

func greatestNumber(n *node) int {
	if n.rNode == nil {
		return n.data
	}
	return greatestNumber(n.rNode)
}

func deleteNode(n *node, nodeValue int) *node {

	if n == nil {
		return nil
	}

	if nodeValue < n.data {
		n.lNode = deleteNode(n.lNode, nodeValue)
		return n
	} else if nodeValue > n.data {
		n.rNode = deleteNode(n.rNode, nodeValue)
		return n
	} else {
		if n.lNode == nil {
			return n.rNode
		}
		if n.rNode == nil {
			return n.lNode
		}

		n.rNode = lift(n.rNode, n)
		return n
	}
}

func lift(node, nodeToDelete *node) *node {
	if node.lNode != nil {
		node.lNode = lift(node.lNode, nodeToDelete)
		return node
	} else {
		nodeToDelete.data = node.data
		return node.rNode
	}
}

func preOrderTraversal(n *node) {
	if n == nil {
		return
	}
	fmt.Printf("%4d", n.data)
	preOrderTraversal(n.lNode)
	preOrderTraversal(n.rNode)
}

func postOrderTraversal(n *node) {
	if n == nil {
		return
	}
	postOrderTraversal(n.lNode)
	postOrderTraversal(n.rNode)
	fmt.Printf("%4d", n.data)
}

func inOrderTraversal(n *node) {
	if n == nil {
		return
	}
	inOrderTraversal(n.lNode)
	fmt.Printf("%4d", n.data)
	inOrderTraversal(n.rNode)
}

func search(n *node, searchValue int) *node {
	if n == nil || n.data == searchValue {
		return n
	}
	if searchValue < n.data {
		return search(n.lNode, searchValue)
	} else {
		return search(n.rNode, searchValue)
	}
}

func insert(parentNode *node, nodeValue int) {
	if nodeValue < parentNode.data {
		if parentNode.lNode == nil {
			parentNode.lNode = &node{data: nodeValue}
			return
		}
		insert(parentNode.lNode, nodeValue)
	} else {
		if parentNode.rNode == nil {
			parentNode.rNode = &node{data: nodeValue}
			return
		}
		insert(parentNode.rNode, nodeValue)
	}
}
