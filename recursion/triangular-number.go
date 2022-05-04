package main

func triangularNumber(num int) int {
	if num == 1 {
		return 1
	}

	return num + triangularNumber(num-1)
}
