package main

func missingAlphabet(str string) rune {

	strMap := make(map[rune]bool, len(str))

	for _, r := range str {
		strMap[r] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		if !strMap[i] {
			return i
		}
	}

	return 0
}
