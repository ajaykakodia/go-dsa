package main

import "fmt"

type LinkedList struct {
	head *node
}

func (l *LinkedList) First() *node {
	return l.head
}

func (l *LinkedList) LastElement() int {

	if l.head == nil {
		fmt.Println("No Item in Linked List")
		return -1
	}

	currentNode := l.head

	for currentNode.NextNode() != nil {
		currentNode = currentNode.NextNode()
	}
	return currentNode.data
}

func (l *LinkedList) getNode(index int) *node {

	if l.head == nil {
		fmt.Println("No Node to return")
		return nil
	}

	currentPosition := 0
	currentNode := l.head

	for currentPosition < index && currentNode != nil {
		currentPosition++
		currentNode = currentNode.NextNode()
	}

	return currentNode
}

func (l *LinkedList) RemoveSpecificNode(n *node) {

	if n == nil && n.NextNode() == nil {
		fmt.Println("No way we can delete this node")
		return
	}

	nextNode := n.NextNode()

	n.data = nextNode.data

	n.next = nextNode.NextNode()

}

func (l *LinkedList) ReverseLinkList() {

	if l.head == nil {
		fmt.Println("No Item to reverse")
		return
	}

	var previousNode, nextNode *node
	currentNode := l.head

	for currentNode != nil {
		nextNode = currentNode.NextNode()
		currentNode.next = previousNode

		previousNode = currentNode
		currentNode = nextNode
	}
	l.head = previousNode
}

func (l *LinkedList) Read() {
	if l.head == nil {
		fmt.Println("No Item in Linked List")
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println("List Item :", currentNode.data)
		currentNode = currentNode.NextNode()
	}
}

func (l *LinkedList) Push(val int) {
	newNode := &node{data: val}
	if l.head == nil {
		l.head = newNode
		return
	}
	currentNode := l.head
	for currentNode.NextNode() != nil {
		currentNode = currentNode.NextNode()
	}

	currentNode.next = newNode
}

func (l *LinkedList) Search(val int) int {
	if l.head == nil {
		return -1
	}
	currentPosition := 0
	currentNode := l.head

	for currentNode != nil {
		if currentNode.data == val {
			return currentPosition
		}
		currentPosition++
		currentNode = currentNode.NextNode()
	}

	return -1
}

func (l *LinkedList) PushAtIndex(index, val int) {
	newNode := &node{data: val}
	if index == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}
	currentIndex := 0
	currentNode := l.head

	for currentIndex < index-1 && currentNode.next != nil {
		currentIndex++
		currentNode = currentNode.NextNode()
	}

	newNode.next = currentNode.NextNode()
	currentNode.next = newNode
}

func (l *LinkedList) RemoveFromIndex(index int) {
	if index == 0 {
		l.head = l.head.NextNode()
		return
	}
	currentIndex := 0
	currentNode := l.head
	for currentIndex < index-1 && currentNode != nil {
		currentIndex++
		currentNode = currentNode.NextNode()
	}

	if currentNode == nil || currentNode.NextNode() == nil {
		fmt.Println("No Node to delete for Index :", index)
		return
	}

	currentNode.next = currentNode.NextNode().next
}

type node struct {
	next *node
	data int
}

func (n *node) NextNode() *node {
	return n.next
}
