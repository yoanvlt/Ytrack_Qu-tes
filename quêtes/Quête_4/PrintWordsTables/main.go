package main
import "github.com/01-edu/z01"
func main(){
	a := []string{"Hello ", "how ", "are ", "you?"}
	PrintWordsTables(a)
}

func PrintWordsTables(a []string){
	for _,mot := range a{
		
		Printrunes(mot)
		z01.PrintRune('\n')
	}
}

func Printrunes(str string){
	for _,m := range str {
		z01.PrintRune(m)
	}
}
