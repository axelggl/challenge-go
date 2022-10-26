package piscine

func Index(s string, toFind string) int {
    index := -1
    s_array := []rune(s)
    f_array := []rune(toFind)
    if len(f_array) == 0 {
        return 0
    } else if len(s_array) < len(f_array) {
        return -1
    } else {
        for i := 0; i < len(s_array); i++ {
            if s_array[i] == f_array[0] {
                index = i
                if len(f_array) == 1 {
                    return index
                }
                for j := 1; j < len(f_array) && j+i < len(s_array); j++ { // si j[1]
                    // tant que j est < a la len find, et j + i < len(s_array) ?
                    if s_array[i+j] == f_array[len(f_array)-1] { // dernière ligne
                        return index
                    }
                    if s_array[i+j] != f_array[j] {
                        return -1
                    }
                }
            }
        }
    }
    return index
}
