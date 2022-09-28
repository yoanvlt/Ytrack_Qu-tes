package main

import "fmt"

func main (){
	n := 9
	fmt.Println(recursivefact(n))

}
func recursivefact (n int) int{
	if n == 1 {
		return 1
	}else{
		return n*recursivefact(n-1)
	}
}
