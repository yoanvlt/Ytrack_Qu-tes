package main

import "fmt"

func main(){
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))
}

func IsLower(s string)bool{
	counter := 0
	for _,lettre := range s {
		if lettre >= 'a' && lettre <='z'{
			counter+=1
		}
	}
	if counter == len(s){
		return true
	}
	return false
}