package piscine

import "fmt"

func Raid1e(x,y int){

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
			}else if y==1 && x!=1 {
				for j := 1; j <= x; j++ {
					if j==1 {
						fmt.Print("A")
					}else if j==y {
						fmt.Print("C")
					}else if j!=1 && j!=x {
						fmt.Print("B")
					}else{
						fmt.Print("C")
					}
				}
				fmt.Print("\n")
				break
			}else{
			//ligne y
			for j := 1; j <= x; j++ {
				//colonne x
				if (i==1 && j==1) || (i==y && j==x)  {
					fmt.Print("A")
				}else if (i==1&&j==x) || (i==y && j==1) {
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
}