package main

import "fmt"

func main(){
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}

func IsUpper(s string)bool {
	counter := 0
	for _,lettre := range s{
		if lettre >= 'A' && lettre <= 'Z'{
			counter++
		}
	}
	if counter == len(s){
		return true
	}
	return false
}
