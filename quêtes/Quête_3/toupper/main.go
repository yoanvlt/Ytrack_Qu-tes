package main

import "fmt"

func main(){
	fmt.Println(ToUpper("Hello! How are you?"))
}

func ToUpper(s string)string {
	var mot string
	for _,lettre := range s {
		if lettre >= 'a' && lettre <= 'z'{	
			mot += string(lettre-32)
		}else{
			mot += string(lettre)
		}
	}
	
	return mot
}