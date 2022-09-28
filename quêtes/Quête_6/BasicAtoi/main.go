package main
import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string)int{
	var res int
	var i int
	for i < len(s){
		if s[i] >= '1' && s[i] <= '9'{
			res = res*10 + int(s[i]-'0')
		}
		i++
	}
	return res
}


