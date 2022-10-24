package piscine

func Index(s string, toFind string) int {
	var index int
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			index = i
			if len(toFind) == 1 {
				return index
			}
			for j := 1; j < len(toFind) && j+i < len(s); j++ {
				if s[i] != toFind[j] {
					return -1
				} else if s[i] == toFind[len(toFind)-1] {
					return index
				}
			}
		}
	}
	return index
}
