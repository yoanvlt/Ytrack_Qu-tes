package main

import (
	"fmt"
)

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(s []string)[]string{
	for i := 0 ; i  < len(s) - 1 ; i++{
		for index,_ := range s{
			if index == len(s)-1{
				break
			}else if s[index] > s[index+1]{
				s[index],s[index+1] = s[index+1],s[index]
			}
		}
	}
	return 	s
}	