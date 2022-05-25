package main

import "fmt"

func main() {
	fmt.Println("Message from Graph....")
	alice := NewVertices("Alice")
	bob := NewVertices("Bob")
	ajay := NewVertices("Ajay")
	candy := NewVertices("Candy")
	darek := NewVertices("Darek")
	fred := NewVertices("Fred")
	halen := NewVertices("Halen")
	gina := NewVertices("Gina")
	irena := NewVertices("Irena")

	alice.AddAdjacentVertex(bob)
	alice.AddAdjacentVertex(candy)
	alice.AddAdjacentVertex(darek)
	alice.AddAdjacentVertex(ajay)

	ajay.AddAdjacentVertex(darek)
	ajay.AddAdjacentVertex(alice)

	bob.AddAdjacentVertex(alice)
	bob.AddAdjacentVertex(fred)

	candy.AddAdjacentVertex(alice)
	candy.AddAdjacentVertex(halen)

	darek.AddAdjacentVertex(alice)
	darek.AddAdjacentVertex(gina)
	darek.AddAdjacentVertex(ajay)

	fred.AddAdjacentVertex(bob)
	fred.AddAdjacentVertex(halen)

	halen.AddAdjacentVertex(fred)
	halen.AddAdjacentVertex(candy)

	gina.AddAdjacentVertex(irena)
	gina.AddAdjacentVertex(darek)

	irena.AddAdjacentVertex(gina)

	DepthTraverse(alice, make(map[string]bool))

	searchVertex := DepthFirstSearch(alice, "Darek", make(map[string]bool))
	fmt.Println("Searched Vertex :", *searchVertex)

	BreadthFirstTraverse(alice)

	searchVertex = BreadthFirstSearch(alice, "Irena")
	fmt.Println("Searched Vertex By BFS:", *searchVertex)

	fmt.Println("---------------------------- Weighted Graph -----------------------------")

	atlanta := NewWVertices("Atlanta")
	boston := NewWVertices("Boston")
	chicago := NewWVertices("Chicago")
	denver := NewWVertices("Denver")
	elPaso := NewWVertices("El Paso")

	atlanta.AddAdjacentVertex(boston, 100)
	atlanta.AddAdjacentVertex(denver, 160)

	boston.AddAdjacentVertex(chicago, 120)
	boston.AddAdjacentVertex(denver, 180)

	chicago.AddAdjacentVertex(elPaso, 80)

	denver.AddAdjacentVertex(chicago, 40)
	denver.AddAdjacentVertex(elPaso, 140)

	elPaso.AddAdjacentVertex(boston, 100)

	DepthTraverseWVertex(atlanta, make(map[string]bool))

	cheapestFare := CheaperFare(atlanta, elPaso)
	fmt.Println("Cheapest Fare: ", cheapestFare)
}
