package main

func GameWinner(numberOfCoin int, currentPlayer string) string {

	if numberOfCoin <= 0 {
		return currentPlayer
	}

	var nextPlayer string

	if currentPlayer == "you" {
		nextPlayer = "them"
	} else {
		nextPlayer = "you"
	}

	if GameWinner(numberOfCoin-1, nextPlayer) == currentPlayer ||
		GameWinner(numberOfCoin-2, nextPlayer) == currentPlayer {
		return currentPlayer
	} else {
		return nextPlayer
	}
}

func GameWinner2(numberOfCoin int) string {

	if (numberOfCoin-1)%3 == 0 {
		return "them"
	}

	return "you"
}
