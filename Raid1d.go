package piscine

import "fmt"

func Raid1d(x,y int){

	if x>0 && y>0 {
		for i := 1; i <= y; i++ {
			//ligne y
			for j := 1; j <= x; j++ {
				//colonne x
				if (i==1 && j==1) || (i==y && j==1)  {
					fmt.Print("A")
				}else if (i==1&&j==x) ||  (i==y && j==x){
					fmt.Print("C")
				}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
					fmt.Print("B")
				}else if (i!=1 && j==1) {
					fmt.Print("B")
				}else if (i!=y && j==x){
					fmt.Print("B")
				}else{
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}