package piscine

func BasicJoin(elems []string) string {
	var final_string string   // revient à faire final_string := ""
	for _, s := range elems { // taille tableau
		final_string += s
	}
	return final_string
}
