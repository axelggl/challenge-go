package piscine

func Join(strs []string, sep string) string {
	var final_string string
	for i, s := range strs {  // s récupère la valeur d'une case dans un tableau
		final_string += s     // +- rajoute à la fin de
		if i != len(strs)-1 { // permet de ne pas mettre de ":" à la fin
			final_string += sep
		}
	}
	return final_string
}
