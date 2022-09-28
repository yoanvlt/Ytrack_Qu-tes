package main

import ("fmt"
	
)

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
	
}

func Any(f func(string) bool, a []string)bool{
	var tab []bool
	for index,_ := range a {
		s := a[index]
		tab = append(tab,f(s))
		if tab[index] == true {
			return true
		}
	}
	return false
}



func IsNumeric(s string)bool{
	counter := 0
	for _,lettre := range s{
		if lettre >= '0' && lettre <= '9'{
			counter+=1
		}
	}
	if counter == len(s){
		return true
	}
	return false
}