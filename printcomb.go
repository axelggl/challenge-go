package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i <= 57; i++ {
		a := rune(i)

		for j := 48; j <= 57; j++ {
			b := rune(j)

			for k := 48; k <= 57; k++ {
				c := rune(k)

				if i < j && j < k {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if i == 55 && j == 56 && k == 57 {
						z01.PrintRune(10)
					} else {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}

/*
   i = 0
   j = 0
   k = 0

   for i = 0; i <=9; i++
   	for j = 0; j <=9; j++
   		for k = 0; k >=9; k++
   			if i > j && j > k
   			print i j k



   || = ou ; && = et


*/
