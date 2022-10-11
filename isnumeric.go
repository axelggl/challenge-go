package piscine

func IsNumeric(s string) bool {
	my_array := []rune(s)
	a := true
	for _, r := range my_array {
		if !(r >= 48 && r <= 57) {
			a = false
		}
	}
	return a
}
