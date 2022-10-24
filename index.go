package piscine

func Index(s string, toFind string) int {
	var index int
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			index = i
			if len(toFind) == 1 {
				return index
			}
			for j := 1; (j < len(toFind) && j+i < len(s)); j++ {
				if s[i] != toFind[j] {
					return -1
				} else if s[i] == toFind[len(toFind) - 1] {
					return index
				}
			}
		}
	}
	return index
}





/*func Index(s string, toFind string) int {
	sTable := []rune(s)
	fTable := []rune(toFind)

	if len(s) < 1 {
		return 0
	} else if len(s) < len(toFind) {
		return -1
	}
	for i := 0; i < len(sTable); i++ {
		if sTable[i] == fTable[0] {
			for j := 0; j < len(toFind); j++ {
				if sTable[i+j] == fTable[j] {
					return i
				} else if sTable[i+j] != fTable[j] {
					return i
				}
			}
		}
	}
	return -1
}
*/