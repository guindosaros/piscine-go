package piscine

import ( 
        
        "strconv"

)

func BasicAtoi2(s string) int {
            a := 0
           intI, err := strconv.Atoi(s)
	if err == nil{
		 a = intI 
	}

return a
}