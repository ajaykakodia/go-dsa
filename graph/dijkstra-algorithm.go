package main

import "fmt"

func CheaperFare(startingCity *WVertex, destinationVertex *WVertex) int {
	visitedCities := make(map[string]bool)
	unvisitedCities := make([]*WVertex, 0)

	cheapestPriceTable := make(map[string]int)
	cheapestPreviousStopoverCity := make(map[string]string)

	currentCity := startingCity

	cheapestPriceTable[startingCity.value] = 0

	for currentCity != nil {
		visitedCities[currentCity.value] = true
		unvisitedCities = deleteFromSlice(unvisitedCities, currentCity.value)

		for adVertex, wight := range currentCity.adjacent_vertices {
			if !visitedCities[adVertex.value] {
				unvisitedCities = append(unvisitedCities, adVertex)
				chepestPriceForAdjacentVertex := cheapestPriceTable[currentCity.value] + wight
				if currentCheapestPrice, okk := cheapestPriceTable[adVertex.value]; !okk || chepestPriceForAdjacentVertex < currentCheapestPrice {
					cheapestPriceTable[adVertex.value] = chepestPriceForAdjacentVertex
					cheapestPreviousStopoverCity[adVertex.value] = currentCity.value
				}
			}
		}

		currentCity = getCheapestUnvisitedCity(unvisitedCities, cheapestPriceTable)
	}

	fmt.Println("Sortest Path for City ", destinationVertex.value)

	currentTravelingCity := destinationVertex.value
	fmt.Printf("%15s", destinationVertex.value)
	for currentTravelingCity != startingCity.value {
		fmt.Printf(" <- %10s", cheapestPreviousStopoverCity[currentTravelingCity])
		currentTravelingCity = cheapestPreviousStopoverCity[currentTravelingCity]
	}

	fmt.Printf("\n%v\n", cheapestPriceTable)
	fmt.Printf("\n%v\n", cheapestPreviousStopoverCity)

	return cheapestPriceTable[destinationVertex.value]
}

func getCheapestUnvisitedCity(sl []*WVertex, cpt map[string]int) *WVertex {
	if len(sl) == 0 {
		return nil
	}

	var cheapestCity *WVertex
	cheapestCityPrice := 0

	for _, vertex := range sl {
		if cheapestCityPrice == 0 || cheapestCityPrice > cpt[vertex.value] {
			cheapestCity = vertex
			cheapestCityPrice = cpt[vertex.value]
		}
	}
	return cheapestCity
}

func deleteFromSlice(sl []*WVertex, val string) []*WVertex {
	valIndex := -1
	for inx, str := range sl {
		if str.value == val {
			valIndex = inx
			break
		}
	}
	if valIndex == -1 {
		return sl
	}

	return append(sl[0:valIndex], sl[valIndex+1:]...)
}
