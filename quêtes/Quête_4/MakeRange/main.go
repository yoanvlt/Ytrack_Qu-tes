package main
import "fmt"
func main(){
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

func MakeRange(min int,max int)[]int{
	size := max - min
	if size > 0 {
		reponse :=  make([]int, size)
		for i := 0; i < size ; i++ {
			reponse[i] = i + min
		}
		return reponse		
	}
	var reponse []int
	return reponse	
}
