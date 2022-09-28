package main

import "fmt"

func main() {
	fmt.Println(RecursivePower(4,4))
}

func  RecursivePower(nb int, nbpower int) int{
	if nbpower == 0 {
		return 1
	}
	return nb *  RecursivePower(nb,nbpower - 1)
}