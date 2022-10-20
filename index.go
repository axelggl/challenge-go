package piscine

func Index(s string, toFind string) int {
        sTable := []rune(s)
        fTable := []rune(toFind)

    if len(s) < 1 {
        return 0
    } else if len(s) < len(toFind) {
        return -1
    }
    for i := 0; i < len(sTable); i++ {
        if sTable[i] == fTable[0] {
            for j := 0; j < len(toFind); j++ {
                if sTable[i+1] == fTable[j] {
                    return i
                } else if sTable[i+1] != fTable[j] {
                    return i
                }
            }
        }
    }
    return -1
}
