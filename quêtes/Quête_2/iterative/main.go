package main

import "fmt"

func main () {
	fmt.Println(iterative(3,4))
}
func iterative (nbpower int,nb int)int {
	res := 1
	for i := 1 ; i <= nbpower ; i++ {
		res = res*nb
	}
	return res
}	