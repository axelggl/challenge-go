package piscine

func Join(strs []string, sep string) string {
	var final_string string
	for _, s := range strs {
		final_string += s
		final_string += sep
	}
	return final_string
}
