package main

import "github.com/01-edu/z01"

func main() {
}

 

func Printnbr(n int){
	if n < 0  {
		z01.PrintRune('-')
		n = -n
	}
	if n > 9  {
		Printnbr(n/10)
	}
	z01.PrintRune(rune(n%10 + 48))
	 
}