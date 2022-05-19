package main

import "fmt"

func main() {
	// go run ./linked-list

	fmt.Println("Calling from linked-list")

	queue := Queue{}
	queue.NewQueue()

	queue.Enqueue(21)
	queue.Enqueue(22)
	queue.Enqueue(23)
	queue.Enqueue(24)
	queue.ReadQueue()
	fmt.Println("______________")
	queue.Dequeue()
	queue.Dequeue()
	queue.ReadQueue()
	fmt.Println("______________")
	queue.Dequeue()
	queue.Dequeue()

	queue.ReadQueue()
	fmt.Println("______________")
	queue.Dequeue()

	dll := DoublyLinkedList{}

	dll.PushAtIndex(0, 20)
	dll.PushAtIndex(0, 21)
	dll.PushAtIndex(0, 22)
	dll.ReverseRead()
	dll.PushAtIndex(0, 19)
	fmt.Println("______________")
	dll.ReverseRead()
	dll.RemoveFromIndex(0)
	fmt.Println("______________")
	dll.ReverseRead()
	dll.RemoveFromIndex(2)
	fmt.Println("______________")
	dll.ReverseRead()

	fmt.Println("______________")

	ll := LinkedList{}
	ll.Push(23)
	ll.Push(50)
	ll.Push(60)
	ll.Read()
	fmt.Println("Last Element in Linked List :", ll.LastElement())
	fmt.Println("______________")
	ll.PushAtIndex(0, 21)
	ll.Read()
	fmt.Println("______________")
	ll.PushAtIndex(3, 70)
	ll.Read()
	fmt.Println("______________")
	ll.PushAtIndex(200, 160)
	ll.Read()
	fmt.Println("______________")
	ll.RemoveFromIndex(3)
	ll.Read()
	fmt.Println("______________")
	ll.RemoveFromIndex(0)
	ll.Read()
	fmt.Println("______________")
	ll.RemoveFromIndex(5)
	ll.Read()
	fmt.Println("______________")
	ll.RemoveFromIndex(3)
	ll.Read()
	fmt.Println("______________")
	index := ll.Search(50)

	if index == -1 {
		fmt.Println("Data not exists in linked list")
	} else {
		fmt.Println("Data found at position :", index+1)
	}

	ll.Push(23)
	ll.Push(51)
	ll.Push(60)
	ll.Push(52)
	ll.Push(53)
	ll.Push(54)
	fmt.Println("______________")
	ll.Read()
	fmt.Println("______________")
	ll.ReverseLinkList()
	ll.Read()

	node := ll.getNode(3)
	fmt.Println("fetched node: ", node.data)
	ll.RemoveSpecificNode(node)
	ll.Read()
}
