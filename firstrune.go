package piscine

func FirstRune(s string) rune {
	s_runes := []rune(s) // crée un tableau avec s au format rune en ASCII
	return s_runes[0]
}
