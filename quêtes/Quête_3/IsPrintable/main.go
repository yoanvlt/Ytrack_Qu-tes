package main
import "fmt"
func main(){
	fmt.Println(IsPrintable("Hello\n"))
	fmt.Println(IsPrintable("01,02,03"))
}

func IsPrintable(s string)bool{
	for _,lettre := range s{
		if lettre < 32 || lettre > 126{
			return false
		}
	}
	return true
}
