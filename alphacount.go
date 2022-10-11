package piscine

func AlphaCount(s string) int {
	my_array := []rune(s)
	k := 0
	for _, r := range my_array {
		if r >= 65 && r <= 90 && r >= 97 && r <= 122{
			k++
		}
	}
	return k
}
