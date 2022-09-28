package main
import "fmt"
func main(){

	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

func SplitWhiteSpaces(s string) []string {
	var res []string
	mot := ""
	for _,lettre := range s{
		if string(lettre) != " "{
			mot += string(lettre)
		}else{
			res = append(res,mot)
			mot = ""
		}
	}
	if mot != ""{
		res = append(res,mot)
	}
	return res
}

