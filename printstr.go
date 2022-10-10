package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	s_runes := []rune(s)
	index := len(s_runes)
	for i := 0; i < (index); i++ {
		z01.PrintRune(s_runes[i])
	}
}
