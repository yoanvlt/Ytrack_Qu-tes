package main

import "github.com/01-edu/z01"

func main () {	
	isNegative(1)
	isNegative(0)
	isNegative(-1)
}

func isNegative (a int) {
	if a >= 0 {
		z01.PrintRune('F') 
	}else {
		z01.PrintRune('T')
	}
	z01.PrintRune('\n')
}