package main

import "fmt"

type PrintManager struct {
	queue QueueStr
}

func (pr *PrintManager) AddInPrintQueue(page string) {
	pr.queue.Enqueue(page)
}

func (pr *PrintManager) Run() {
	for pr.queue.Size() > 0 {
		pageToPrint, _ := pr.queue.Dequeue()
		fmt.Println("Printing Page: ", pageToPrint)
	}
}
