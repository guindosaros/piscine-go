package piscine

import "fmt"

func Raid1c(x,y int){

	if x>0 && y>0 {

		for i := 1; i <= y; i++ {
			if x==1 && y!=1{
				for j := 1; j <= y; j++ {

					if j==1 {
						fmt.Print("A")
					}else if j==y {
						fmt.Print("C")
					}else if j!=1 && j!=y {
						fmt.Print("B")
					}
					fmt.Print("\n")
				}
				break
			}else{
			//ligne y
			for j := 1; j <= x; j++ {
				//colonne x
				
				if (i==1 && j==1) || (i==1 && x==j) {
					fmt.Print("A")
				}else if (i==y && j==1) || (i==y && j==x) {
					fmt.Print("C")
				}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=1 || j!=x)) || (j==1 && (i!=1 || i==y)) || (j==x && (i!=1 || i!=y)){
					fmt.Print("B")
				}else{
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
}