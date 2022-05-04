package main

func uniquePath(row, col int) int {
	if row == 1 || col == 1 {
		return 1
	}

	return uniquePath(col-1, row) + uniquePath(col, row-1)
}
