package main

type Vertex struct {
	value             string
	adjacent_vertices []*Vertex
}

func NewVertices(val string) *Vertex {
	return &Vertex{value: val, adjacent_vertices: make([]*Vertex, 0)}
}

func (v *Vertex) AddAdjacentVertex(vertex *Vertex) {
	v.adjacent_vertices = append(v.adjacent_vertices, vertex)
}

type WVertex struct {
	value             string
	adjacent_vertices map[*WVertex]int
}

func NewWVertices(val string) *WVertex {
	return &WVertex{value: val, adjacent_vertices: make(map[*WVertex]int)}
}

func (v *WVertex) AddAdjacentVertex(vertex *WVertex, val int) {
	v.adjacent_vertices[vertex] = val
}
