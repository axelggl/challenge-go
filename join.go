package piscine

func Join(strs []string, sep string) string {
	var final_string string
	for i, s := range strs {
		final_string += s
		if i != len(strs)-1 {
			final_string += sep
		}
	}
	return final_string
}
