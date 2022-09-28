package main

import "fmt"

func main(){
	fmt.Println(AlphaCount("Hello 78 World!    4455 /"))
}
func AlphaCount(s string)int {
	counter := 0
	for _, lettre := range s{
		if lettre >= 'a' && lettre <='z'  {
			counter++
		}else if lettre >= 'A' && lettre <= 'Z'{
			counter++
		}
	}
	return counter
}