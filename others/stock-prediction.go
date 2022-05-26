package main

func StockPrediction(arr []int) int {
	maxprofit := 0
	currentProfit := 0
	lowestPurchase := arr[0]

	for i := 0; i < len(arr); i++ {
		currentProfit = arr[i] - lowestPurchase
		if currentProfit > maxprofit {
			maxprofit = currentProfit
		}
		if lowestPurchase > arr[i] {
			lowestPurchase = arr[i]
		}
	}

	return maxprofit
}
