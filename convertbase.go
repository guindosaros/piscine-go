package piscine

import (
	"math"
	"strings"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return  NbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func NbrBase(nbr int, str string) string {
	indx := 0
	for _,res:= range str {
		if string(res) == "-" || string(res) == "+" || strings.Count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2{
		return "0"
	}else if math.MaxInt32 <nbr || math.MinInt32 > nbr{
		return string(nbr)
	}else{
		neg := ""
		if nbr < 0 {
			neg = "-"
			nbr *= -1	
		}
		i:=0
		nan:= ""
		for nbr >= len(str) {
			if nbr >= len(str) {
				nan +=string(str[nbr % len(str)])
				nbr = nbr/len(str)
				i++
			}
		}
		nan +=string(str[nbr])
		return neg+Reverse(nan)
	}
}


