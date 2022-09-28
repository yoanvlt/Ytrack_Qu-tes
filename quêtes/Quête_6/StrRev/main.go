package main
import "fmt"

func main() {
	s := "Hello World!"
	fmt.Println(StrRev(s))
}

func StrRev(s string)string{
	var tab []rune 
	for i := len(s)-1 ; i >s= 0 ; i--{
		tab = append(tab,rune(s[i]))
	}	
	return string(tab)
}	
