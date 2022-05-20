package main

import "fmt"

type BinaryHeap struct {
	data []int
}

func (h *BinaryHeap) rootNode() int {
	if h.data != nil {
		return h.data[0]
	}
	return -1
}

func (h *BinaryHeap) lastNode() int {
	if h.data != nil {
		return h.data[len(h.data)-1]
	}
	return -1
}

func parentNodeIndex(index int) int {
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return index*2 + 1
}

func rightChildIndex(index int) int {
	return index*2 + 2
}

func (h *BinaryHeap) InsertNode(val int) {

	h.data = append(h.data, val)

	lastIndex := len(h.data) - 1
	var parentIndex int
	for {
		parentIndex = parentNodeIndex(lastIndex)
		if val > h.data[parentIndex] {
			h.data[lastIndex], h.data[parentIndex] = h.data[parentIndex], h.data[lastIndex]
			lastIndex = parentIndex
			continue
		}
		return
	}
}

func (h *BinaryHeap) hasGreaterChild(index int) bool {
	lastNodeIndex := len(h.data) - 1
	lci := leftChildIndex(index)

	if lci > lastNodeIndex {
		return false
	}

	if h.data[lci] > h.data[index] {
		return true
	}

	rci := rightChildIndex(index)

	if rci <= lastNodeIndex && h.data[rci] > h.data[index] {
		return true
	}

	return false
}

func (h *BinaryHeap) calculateLargerChildIndex(index int) int {
	lastNodeIndex := len(h.data) - 1
	lci := leftChildIndex(index)
	rci := rightChildIndex(index)

	if rci > lastNodeIndex {
		return lci
	}

	if h.data[lci] > h.data[rci] {
		return lci
	}
	return rci
}

func (h *BinaryHeap) Delete() int {

	if len(h.data) == 0 {
		fmt.Println("No data to remove, heap is empty")
		return -1
	}
	deletedData := h.data[0]
	h.data[0] = h.data[len(h.data)-1]

	h.data = h.data[:len(h.data)-1]
	index := 0

	for h.hasGreaterChild(index) {
		gci := h.calculateLargerChildIndex(index)
		h.data[gci], h.data[index] = h.data[index], h.data[gci]
		index = gci
	}
	return deletedData
}
