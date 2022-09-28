package main

import "fmt"

func main(){
	fmt.Println(Concat("Hello!", " How are you?"))
}

func Concat(str1,str2 string)string{
	str1 +=str2
	return str1
}