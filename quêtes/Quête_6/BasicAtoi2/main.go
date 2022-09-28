package main
import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string)int{
	var res int
	for i,_ := range s{
		if s[i] >= '0' && s[i] <= '9'{
			if s[i] != '0'{
				res = res*10 + int(s[i]-'0')
			}
		}else {
			return 0
		}	
	}
	return res
}