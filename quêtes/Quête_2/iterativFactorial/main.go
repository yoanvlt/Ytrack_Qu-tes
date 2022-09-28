package main

import "fmt"
func main (){
	n := 9
	fmt.Println(iteratibeFactorial(n))
}

func iteratibeFactorial (n int) int{
	res := 1	
	for i := 1 ; i <= n ; i++ {
		res = res * i
		}
	return res
}