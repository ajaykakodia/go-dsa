package main

func findNeedleInHayStack(haystack, needle string) bool {
	needleFound := false

	for i := 0; i < len(haystack); i++ {
		a := i
		if haystack[i] == needle[0] && (i+len(needle) < len(haystack)) {
			a += 1
			needleFound = true
			for j := 1; j < len(needle); j++ {
				if haystack[a] != needle[j] {
					needleFound = false
					break
				}
				a += 1
			}
		}
		if needleFound {
			return needleFound
		}
	}
	return needleFound
}

func findNeedleInHayStack2(haystack, needle string) bool {
	needleFound := false
	nl := len(needle)
	hl := len(haystack)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && (i+nl < hl) && haystack[i:i+nl] == needle {
			return true
		}
	}
	return needleFound
}
