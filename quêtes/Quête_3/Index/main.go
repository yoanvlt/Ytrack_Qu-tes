package main

import "fmt"

func main(){
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string,str string)int{
	for index,lettre := range s{
		if lettre == rune(str[0]){
			for index2,lettre2 := range str{
				if lettre2 != rune(s[index+index2]){
					break
				}else if index2 == len(str)-1{
					return index
				}

			}

		} 
	}
	return -1
}