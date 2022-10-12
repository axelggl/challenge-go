package piscine

func ToLower(s string) string {
	my_sentence := []rune(s)
	for i := 0; i < len(my_sentence); i++ {
		if my_sentence[i] >= 65 && my_sentence[i] <= 90 {
			my_sentence[i] += 32
		}
	}
	return string(my_sentence)
}
