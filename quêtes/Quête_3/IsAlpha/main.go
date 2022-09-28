package main
import "fmt"
func main(){
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}

func IsAlpha(s string)bool{
	counter := 0
	for _,lettre := range s{
		if lettre >='a' && lettre <= 'z' || lettre >='A' && lettre <= 'Z' || lettre >= '0' && lettre <= '9'{
			counter+=1
		}
	}
	if counter == len(s){
		return true
	}
	return false
}