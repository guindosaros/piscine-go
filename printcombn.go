package piscine

import "fmt"
import "strconv"

func printi(b int, n int, a int, i int,l string)  {
	a++
	for j:= b+1; j<10-n+a; j++{
		if(a<n){
			printi(j,n,a,i,l+strconv.Itoa(j))
		}else{
			fmt.Print(l)
			fmt.Print(j)
			if i<10-n{
				fmt.Print(", ")
			} else{
				fmt.Print("\n")
			}
		}
	}
}

func PrintCombN(n int)  {
	for i:=0; i<10; i++{
		if n>1{
			printi(i,n,1,i,strconv.Itoa(i))
		} else{
			fmt.Print(i)
			if i<10-n{
				fmt.Print(", ")
			} else{
				fmt.Print("\n")
			}
		}
	}
}