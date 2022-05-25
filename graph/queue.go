package main

import "fmt"

type Queue struct {
	arr []Vertex
}

func (q *Queue) Enqueue(val Vertex) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Dequeue() *Vertex {
	ql := len(q.arr)
	if ql == 0 {
		fmt.Println("No data exists to Dequeue")
		return nil
	}

	dqe := q.arr[0:1]
	if ql == 1 {
		q.arr = []Vertex{}
	} else {
		q.arr = q.arr[1:]
	}

	return &dqe[0]
}

func (q *Queue) Size() int {
	return len(q.arr)
}
