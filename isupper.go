package piscine

func IsUpper(s string) bool {
    my_array := []rune(s) // on crée un tableau de rune s
    a := true // a est affecté a true car si on return simplement true il quitte la condition et return simplement au fichier main
    for _, r := range my_array {
        if !(r >= 65 && r <= 90) { // si la case [i] qui s'imcrémente a chaque fois // i = 0 tant que i est < à la longeur du tableau, alors i s'incrémente
            a = false // donc la variable a est affecté a vrai
        }
        // il retourne donc la valeur a
    }
    return a
}
