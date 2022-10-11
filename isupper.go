package piscine

func IsUpper(s string) bool {
	my_array := []rune(s)
	a := true

	for i := 0; i <= len(my_array)-1; i++ {
		if my_array[i] >= 65 && my_array[i] <= 90 {
			a = true
		} else {
			a = false
		}
	}
	return a
}
