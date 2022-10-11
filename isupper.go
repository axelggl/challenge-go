package piscine

func IsUpper(s string) bool {
	my_array := []rune(s)
	a := false

	for i := 0; i <= len(my_array)-1; i++ { // range permet de décomposer le tableau similaire à i:= 0; i < len(my_array); i++
		if my_array[i] >= 65 && my_array[i] <= 90 {
			a = true
		} else {
			a = false
		}
	}

	return a
}
