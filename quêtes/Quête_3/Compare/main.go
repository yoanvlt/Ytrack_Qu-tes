package main

import "fmt"

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}
func Compare(a string,b string) int{
	if len(a) == len(b){
		return 0
	}
	if len(a) > len(b){
		return 1
	}	
	return -1
}	