package main

import "fmt"

type DoublyLinkedList struct {
	head *dlNode
	tail *dlNode
}

func (l *DoublyLinkedList) First() *dlNode {
	return l.head
}

func (l *DoublyLinkedList) Last() *dlNode {
	return l.tail
}

func (l *DoublyLinkedList) Read() {
	if l.head == nil {
		fmt.Println("No Item in Double Linked List")
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println("List Item :", currentNode.data)
		currentNode = currentNode.NextNode()
	}
}

func (l *DoublyLinkedList) ReverseRead() {
	if l.tail == nil {
		fmt.Println("No Item in Double Linked List to Read")
		return
	}
	currentNode := l.tail
	for currentNode != nil {
		fmt.Println("List Item :", currentNode.data)
		currentNode = currentNode.PreviousNode()
	}
}

func (l *DoublyLinkedList) Push(val int) {
	newNode := &dlNode{data: val}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode

	l.tail = newNode
}

func (l *DoublyLinkedList) RemoveFromFront() {
	l.RemoveFromIndex(0)
}

func (l *DoublyLinkedList) Search(val int) int {
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

func (l *DoublyLinkedList) PushAtIndex(index, val int) {
	newNode := &dlNode{data: val}
	if index == 0 {
		if l.head == nil {
			l.head = newNode
			l.tail = newNode
			return
		}
		newNode.next = l.head
		l.head.prev = newNode
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
	currentNode.NextNode().prev = newNode
	newNode.prev = currentNode
	currentNode.next = newNode

	if newNode.NextNode() == nil {
		l.tail = newNode
	}
}

func (l *DoublyLinkedList) RemoveFromIndex(index int) {
	if l.head == nil {
		fmt.Println("No data to remove")
		return
	}

	if index == 0 {
		l.head = l.head.NextNode()
		if l.head == nil {
			l.tail = nil
			return
		}
		l.head.prev = nil
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

	nextEle := currentNode.NextNode().next
	if nextEle != nil {
		nextEle.prev = currentNode
	}
	currentNode.next = nextEle
	if currentNode.NextNode() == nil {
		l.tail = currentNode
	}
}

type dlNode struct {
	next *dlNode
	prev *dlNode
	data int
}

func (n *dlNode) NextNode() *dlNode {
	return n.next
}

func (n *dlNode) PreviousNode() *dlNode {
	return n.prev
}
