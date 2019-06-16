package piscine

import(
	"math"
	"fmt"
	"strings"
)

func Reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}

func PrintNbrBase(nbr int, str string)(){
	indx := 0
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2{
		fmt.Print("NV")	
	}else if math.MaxInt32 <nbr || math.MinInt32 > nbr{
		fmt.Print(int64(nbr))
	}else{
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1	
		}
		i:=0
		nan:=""
		for nbr >= len(str) {
			if nbr >= len(str) {
				nan +=string(str[nbr % len(str)])
				nbr = nbr/len(str)
				i++
			}
		}
		nan +=string(str[nbr])
		fmt.Print(Reverse(nan))
	}
}
