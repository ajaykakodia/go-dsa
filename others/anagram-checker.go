package main

func AnagramChecker(str1, str2 string) bool {

	strMap1, strMap2 := make(map[rune]int), make(map[rune]int)

	for _, rne := range str1 {
		strMap1[rne] += 1
	}

	for _, rne := range str2 {
		strMap2[rne] += 1
	}

	for key, val := range strMap1 {
		if strMap2[key] != val {
			return false
		}
		delete(strMap2, key)
	}

	return len(strMap2) == 0
}
