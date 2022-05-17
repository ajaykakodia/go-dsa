package main

type Queue struct {
	dll *DoublyLinkedList
}

func (q *Queue) NewQueue() {
	q.dll = &DoublyLinkedList{}
}

func (q *Queue) Enqueue(val int) {
	q.dll.Push(val)
}

func (q *Queue) Dequeue() {
	q.dll.RemoveFromFront()
}

func (q *Queue) ReadQueue() {
	q.dll.Read()
}
