package main

func anagramGen(str string) []string {
	col := []string{}
	if len(str) == 1 {
		col = append(col, str)
		return col
	}
	char := str[0]
	angs := anagramGen(str[1:])
	for _, ang := range angs {
		l := len(ang)
		for i := 0; i <= l; i++ {
			if i == l {
				na := append([]byte(ang), char)
				col = append(col, string(na))
				continue
			}
			bs := []byte(ang)
			na := append(bs[:i+1], bs[i:]...)
			na[i] = char
			col = append(col, string(na))
		}

	}

	return col
}
