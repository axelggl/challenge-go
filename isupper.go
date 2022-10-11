package piscine

func IsUpper(s string) bool {
	t := []rune(s)
	a := false
	for _, r := range t {
		if r >= 65 && r <= 90 {
			a = true
		} else {
			a = false
		}
	}
	return a
}