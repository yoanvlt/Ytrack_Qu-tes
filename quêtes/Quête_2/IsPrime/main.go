package main

import "fmt"

func main(){
	fmt.Println(IsPrime(2))
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

func FindNextPrime (n int)int {
	result := n 
	for i := n  ; IsPrime(i)== false ; i++{
		result = i + 1
	}
	return result
}