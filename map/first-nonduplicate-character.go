package main

func firstNonDuplicateCharacterInString(str string) string {

	runeCountMap := make(map[rune]int, 0)

	for _, rn := range str {
		runeCountMap[rn] += 1
	}

	for _, rn := range str {
		if runeCountMap[rn] == 1 {
			return string(rn)
		}
	}
	return ""
}
