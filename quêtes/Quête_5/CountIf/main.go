package main
import "fmt"

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

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}

func CountIf(f func(string) bool, tab []string) int {
	var tri []bool
	counter := 0
	for index,_ := range tab{
		s := tab[index]
		tri = append(tri,f(s))
		if tri[index] == true {
			counter += 1
		}
	}
	return counter
}
