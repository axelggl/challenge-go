package piscine

func IsUpper(s string) bool {
    runeBoard := []rune(s)
    a := true
    for _, r := range runeBoard {
        if r >= 65 && r <= 90 {
            a = true
        } else {
            a = false
        }
    }
    return a
}
