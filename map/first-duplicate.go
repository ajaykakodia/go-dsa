package main

func firstDuplicateString(arr1 []string) string {

	strMap := make(map[string]bool, len(arr1))

	for _, val := range arr1 {
		if strMap[val] {
			return val
		}
		strMap[val] = true
	}

	return ""
}
