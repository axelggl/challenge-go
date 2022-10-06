package main

import "github.com/01-edu/z01"

func main() {
	for i := 97; i <= 122; i++ {
		c := rune(i)
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
