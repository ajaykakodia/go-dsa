package main

func findFirstX(str string) int {
	if str[0] == 'x' {
		return 0
	}

	return 1 + findFirstX(str[1:])
}
