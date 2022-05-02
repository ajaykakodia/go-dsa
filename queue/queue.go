package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	arr []int
}

func (q *Queue) Enqueue(val int) {
	q.arr = append(q.arr, val)
}

func (q *Queue) Dequeue() (int, error) {
	ql := len(q.arr)
	if ql == 0 {
		fmt.Println("No data exists to Dequeue")
		return 0, errors.New("no data exists")
	}

	dqe := q.arr[0:1]
	if ql == 1 {
		q.arr = []int{}
	} else {
		q.arr = q.arr[1:]
	}

	return dqe[0], nil
}

func (q *Queue) Seek() int {
	ql := len(q.arr)
	if ql == 0 {
		fmt.Println("No data exists to Dequeue")
		return 0
	}

	dqe := q.arr[0:1]
	return dqe[0]
}

func (q *Queue) Size() int {
	return len(q.arr)
}

type QueueStr struct {
	arr []string
}

func (q *QueueStr) Enqueue(val string) {
	q.arr = append(q.arr, val)
}

func (q *QueueStr) Dequeue() (string, error) {
	ql := len(q.arr)
	if ql == 0 {
		fmt.Println("No data exists to Dequeue")
		return "", errors.New("no data exists")
	}

	dqe := q.arr[0:1]
	if ql == 1 {
		q.arr = []string{}
	} else {
		q.arr = q.arr[1:]
	}

	return dqe[0], nil
}

func (q *QueueStr) Seek() string {
	ql := len(q.arr)
	if ql == 0 {
		fmt.Println("No data exists to Dequeue")
		return ""
	}

	dqe := q.arr[0:1]
	return dqe[0]
}

func (q *QueueStr) Size() int {
	return len(q.arr)
}
