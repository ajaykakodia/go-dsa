package main

func fibonacciSeries(n int) int {
	countSteps++
	if n == 0 || n == 1 {
		return n
	}
	return fibonacciSeries(n-2) + fibonacciSeries(n-1)
}

func fibonacciSeriesMemoization(n int, ma map[int]int) int {
	countSteps++
	if n == 0 || n == 1 {
		return n
	}
	sle, ok := ma[n-2]
	if !ok {
		sle = fibonacciSeriesMemoization(n-2, ma)
		ma[n-2] = sle
	}

	lae, ok := ma[n-1]
	if !ok {
		lae = fibonacciSeriesMemoization(n-1, ma)
		ma[n-1] = lae
	}
	return sle + lae
}

func fibonacciSeriesBottomUp(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1

	for i := 1; i < n; i++ {
		temp := a
		a = b
		b = b + temp
	}
	return b
}
