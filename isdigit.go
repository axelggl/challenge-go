package piscine

func IsDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}
