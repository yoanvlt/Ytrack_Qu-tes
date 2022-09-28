package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result :=Map(IsPrime, a)
	fmt.Println(result)
	fmt.Println(IsPrime(5))
}

func Map(f func(int) bool, a []int) []bool {
	var res []bool
	for _,n := range a {
		res = append(res,f(n))
	}
	return res
}

func IsPrime(n int) bool {
    if n == 1 {
        return false
    }
    var compteur int = 0
    for i := 2; i < n-1; i++ {
        if n%i == 0 {
            compteur = compteur + 1
        }
    }
    if compteur != 0 {
        return false
    }
    return true
}