package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

func NRune (s string, n int)rune {
	for index,lettre := range s{
		if index == n{
			return lettre
		}
	}
	return 0
}

