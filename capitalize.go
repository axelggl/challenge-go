package piscine

func Capitalize(s string) string {
	array := []rune(s)
	if IsAlphanumeric(s) {
		for i := 0; i < len(s); i++ {
			if !IsAlphanumeric(array[i-1]) || array[i] == 0 {
			}

		}
	}
}


func IsAlphanumeric (s string) bool {
	my_array := []rune(s)
	a := true
	for _, r := range my_array {
		if !(r >= 48 && r <= 57 || r >= 65 && r <= 90 || r >= 97 && r <= 122) {
			a = false
		}
	}
	return a
}

func IsLetter (s string) bool {
	my_array := []rune(s)
	a := true
	for _, r := range my_array {
		if !(r >= 65 && r <= 90 || r >= 97 && r <= 122) {
			a = false
		}
	}
	return a
}


/* l'objectif de cette fonction est de mettre en capitale la 
première lettre de chaque mot, les autres lettres doivent être en 
minuscule.
ATTENTION : un mot est considéré comme une suite alphanumérique
en gros si le premier caractère d'un mot est un chiffre,
si le deuxième est une lettre, elle restera en minuscule
si elle était déjà en minuscule, elle se mettra en minuscule
si elle était en majuscule. 
*/

/* 
si alphanumeric de s est vérifié :
	 si première lettre de s est minuscule :
	 	lettre - 32
*/