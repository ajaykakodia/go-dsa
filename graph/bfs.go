package main

import "fmt"

func BreadthFirstTraverse(firstVertex *Vertex) {
	visitedVertex := make(map[string]bool)
	visitedVertex[firstVertex.value] = true

	queue := Queue{}

	queue.Enqueue(*firstVertex)

	for queue.Size() > 0 {
		currentVertex := queue.Dequeue()
		fmt.Println("Visiting Vertex :", currentVertex.value)

		for _, adVertex := range currentVertex.adjacent_vertices {
			if visitedVertex[adVertex.value] {
				continue
			}
			visitedVertex[adVertex.value] = true
			queue.Enqueue(*adVertex)
		}
	}
}

func BreadthFirstSearch(firstVertex *Vertex, searchValue string) *Vertex {
	if firstVertex.value == searchValue {
		return firstVertex
	}

	visitedVertex := make(map[string]bool)
	visitedVertex[firstVertex.value] = true

	queue := Queue{}

	queue.Enqueue(*firstVertex)

	for queue.Size() > 0 {
		currentVertex := queue.Dequeue()

		for _, adVertex := range currentVertex.adjacent_vertices {
			if visitedVertex[adVertex.value] {
				continue
			}
			if adVertex.value == searchValue {
				return adVertex
			}
			visitedVertex[adVertex.value] = true
			queue.Enqueue(*adVertex)
		}
	}

	return nil
}
