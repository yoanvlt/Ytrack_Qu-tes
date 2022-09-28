package main

import "github.com/01-edu/z01"

func main () {
	printComb()
 
}
	func printComb () {
		for a := '0' ; a <= '9' ; a ++ {
			for b := '1' ; b <= '9' ; b ++{
				for c :=  '2' ; c<= '9' ; c++{
					if a < b && b < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						if a == '7' && b == '8' && c == '9'{ 
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
					
				}
			}

		}	
	}	
	