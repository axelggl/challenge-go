package piscine

func IsUpper(s string) bool {
	my_array := []rune(s)
	a := true
	for _, r := range my_array {
		if !(r >= 65 && r <= 90) {
			a = false
			}
	}
	return a
}
