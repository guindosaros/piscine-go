package piscine

import ( 
        
        "strconv"

)

func BasicAtoi(s string) int {
            a := 0
           intI, err := strconv.Atoi(s)
	if err == nil{
		 a = intI 
	}

return a
}