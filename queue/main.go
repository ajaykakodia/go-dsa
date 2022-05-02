package main

import "fmt"

func main() {
	fmt.Println("main from queue folder")
	que := Queue{}

	que.Enqueue(5)
	que.Enqueue(7)
	que.Enqueue(10)
	que.Enqueue(15)
	que.Enqueue(50)
	que.Enqueue(25)

	fmt.Println(que.arr)
	fmt.Println(que.Seek())
	el, _ := que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)
	el, _ = que.Dequeue()
	fmt.Println(el, que.arr)

	// ********************************** //

	fmt.Println("Print Queue start")

	prm := PrintManager{
		QueueStr{},
	}

	prm.AddInPrintQueue("First Document")
	prm.AddInPrintQueue("Second Document")
	prm.AddInPrintQueue("Third Document")
	prm.AddInPrintQueue("Fourth Document")
	prm.AddInPrintQueue("Fifth Document")

	prm.Run()
}
