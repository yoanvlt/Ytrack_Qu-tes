package main

import "fmt"

func main(){
	fmt.Println(tolower("Hello World"))
}

func tolower(s string)string{
	mot := ""
	for _,lettre := range s{
		if lettre >= 'A' && lettre <= 'Z'{
			mot += string(lettre+32)
		}else {
			mot += string(lettre)
		}
	}
	return mot
}