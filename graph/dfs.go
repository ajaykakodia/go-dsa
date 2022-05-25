package main

import "fmt"

func DepthTraverse(vertex *Vertex, visitedVertex map[string]bool) {
	visitedVertex[vertex.value] = true
	fmt.Println("Visited Vertex: ", vertex.value)
	for _, adVertex := range vertex.adjacent_vertices {
		if !visitedVertex[adVertex.value] {
			DepthTraverse(adVertex, visitedVertex)
		}
	}
}

func DepthTraverseWVertex(vertex *WVertex, visitedVertex map[string]bool) {
	visitedVertex[vertex.value] = true
	fmt.Println("Visited Vertex: ", vertex.value)
	for adVertex := range vertex.adjacent_vertices {
		if !visitedVertex[adVertex.value] {
			DepthTraverseWVertex(adVertex, visitedVertex)
		}
	}
}

func DepthFirstSearch(vertex *Vertex, searchValue string, visitedVertex map[string]bool) *Vertex {

	if searchValue == vertex.value {
		return vertex
	}
	visitedVertex[vertex.value] = true
	fmt.Println("Visited Vertex: ", vertex.value)
	for _, adVertex := range vertex.adjacent_vertices {
		if adVertex.value == searchValue {
			return adVertex
		}
		if !visitedVertex[adVertex.value] {
			DepthFirstSearch(adVertex, searchValue, visitedVertex)
		}
	}
	return nil
}
