package piscine

func Index(s string, toFind string) int {
    my_array := []rune(s)
    find_letter := []rune(toFind)

    var final int
    for i, r := range my_array { // on vient rcupérer notre tableau my_array
        if find_letter[0] == r { // si la première est = r
            for _, k := range find_letter { // regarde le contenue et
                if k == r {
                    return i
                }
            }
        } else if find_letter[0] != r {
            for l, k := range find_letter {
                if k == r {
                    a := l

                    final := a
                    return -final
                }
            }
        }
    }
    return final
}
