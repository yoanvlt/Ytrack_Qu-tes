package main
import "github.com/01-edu/z01"
func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for i,_ := range s{
		z01.PrintRune(rune(s[i]))
	}
}