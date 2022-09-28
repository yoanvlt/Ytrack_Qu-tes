package main
import "fmt"
func main(){
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}

func IsNumeric(s string)bool{
	counter := 0
	for _,lettre := range s{
		if lettre >= '0' && lettre <= '9'{
			counter+=1
		}
	}
	if counter == len(s){
		return true
	}
	return false
}