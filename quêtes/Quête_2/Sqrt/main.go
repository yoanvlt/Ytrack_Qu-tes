package main

import "fmt"

func main(){
	fmt.Println(sqrt(49))
}
func sqrt(n int) int{
	for i := 1 ; n/i >= i ; i++ {
		if n/i == i{
			return i
		}
	}
	return 0
}