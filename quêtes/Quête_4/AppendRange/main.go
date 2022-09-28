package main

import "fmt"

func main(){
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

func AppendRange(min int,max int)[]int{
	var reponse []int
	for index := min; index < max ; index++ {
		reponse = append(reponse,index)
	}
	return reponse
}