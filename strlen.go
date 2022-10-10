package piscine

func StrLen(s string) int {
	s_int := []rune(s)  // tableau créé sous forme de rune
	index := len(s_int) // donne la longueur du tableau
	return index        // retourne la longueur du tableau
}
