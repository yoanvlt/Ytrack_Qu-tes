package main
import "fmt"
func main(){
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
func ConcatParams(str []string) string{
	mot := ""
	for _,p := range str{
		mot+= p
		mot+= "\n"
	}
	return mot
}