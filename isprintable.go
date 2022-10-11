package piscine

func IsPrintable(s string) bool {
	my_array := []rune(s)
	a := true
	for _, r := range my_array {
		if r >= 0 && r <= 31 {
			a = false
		}
	}
	return a
}
