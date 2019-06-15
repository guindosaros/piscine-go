package piscine

import (
	"fmt"
	"strings"
)

func PrintNbrBase(nbr int, str string){
	indx := 0
	max := ""
	
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1{
			indx = 1
			break
		}
		if string(res) > max {
			max = string(res)
		}
	}
	
	check := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ind := strings.Index(check, max) + 1
	if indx == 1 || ind < 2 {
		fmt.Print("NV")	
	}else{
		if nbr < 0 {
			fmt.Print("-")
			nbr *= -1	
		}
		for nbr > ind {
			if nbr > ind {
				fmt.Print(nbr % ind)
				nbr = int(nbr/ind)
			}
			
		}
		fmt.Print(nbr)
	}
}
