package main

import "fmt"

type LinkedList struct {
	head *node
}

func (l *LinkedList) First() *node {
	return l.head
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
