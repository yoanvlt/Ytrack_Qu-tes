package main

import "github.com/01-edu/z01"

func main () {
	printComb2()
 
}

func printComb2 () {
	for c := '0' ; c <= '9' ; c++ {
		for d := '0'; d <= '9' ; d++{
			for a := '0' ; a <= '9' ; a++{
				for b := '0' ; b <= '9' ; b++{
					if a > c || a == c && b > d{
						z01.PrintRune(c)
						z01.PrintRune(d)
						z01.PrintRune(' ')
						z01.PrintRune(a)
						z01.PrintRune(b)
							
						if a == '9' && b == '9' && c == '9' && d == '8'{
							break
							}
							
						z01.PrintRune(',')
						z01.PrintRune(' ')
						
						}

					}			
				}
			}
		}
	}
