package main

import (
	"fmt"
)

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(Trier, a1)
	result2 := IsSorted(Trier, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, tab []int) bool {
	for index,_ := range tab{
		if index == len(tab) -1{
			break
		}
		a := tab[index]
		b := tab[index+1]
		s := f(a,b)
		if s > 0 {
			return false
		}
	}
	return true
}


func Trier(a,b int) int{
	if a > b {
		return 1
	}else if a == b {
		return 0
	}
	return -1
}