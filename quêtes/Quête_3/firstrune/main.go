package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune{
	for index, lettre := range s {
		if index == 0 {
			return lettre
		}
	}
	return 0
}