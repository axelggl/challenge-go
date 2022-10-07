package piscine

func LastRune(s string) rune {
	s_runes := []rune(s)    // [] permet de créer un tableau dénombré de 0 à n où 0 est le premier charactère
	index := len(s_runes)   // len permet de donner la longueur d'un string de 1 à n, ex: "Salut c'est moi" va donner une len de 15 où S est 1, a est 2, etc.
	return s_runes[index-1] // on met index-1 car si on avait mis index tout court on aurait eu une erreur dans le output
} // exemple d'erreur : panic: runtime error: index out of range [6] with length 6
