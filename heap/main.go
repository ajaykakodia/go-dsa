package main

import "fmt"

func main() {
	fmt.Println("Message from Heap folder")

	heap := BinaryHeap{}
	heap.InsertNode(55)
	heap.InsertNode(22)
	heap.InsertNode(34)
	heap.InsertNode(10)
	heap.InsertNode(2)
	heap.InsertNode(99)
	heap.InsertNode(68)
	heap.InsertNode(78)
	heap.InsertNode(101)
	heap.InsertNode(56)

	fmt.Println("Data in array: ", heap.data)

	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
	fmt.Println("Deleted data : ", heap.Delete(), " Array after delete : ", heap.data)
}
