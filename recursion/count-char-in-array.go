package main

func countCharInArray(strs []string) int {
	if len(strs) == 1 {
		return len(strs[0])
	}

	return len(strs[0]) + countCharInArray(strs[1:])
}
