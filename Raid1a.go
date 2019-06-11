package piscine

import "fmt"

func Raid1a(x,y int){

	if x>0 && y>0 {

		for i := 1; i <= y; i++ {
			if x==1 && y!=1 {
			for j := 1; j <= y; j++ {

				if j==1 || j==y {
					fmt.Print("o\n")
				}else if j!=1 && j!=y {
					fmt.Print("|\n")
				}
			}
			break
			}else{
				for j := 1; j <= x; j++ {
					if (i==1 && j==1) || (i==1&&j==x) || (i==y && j==1) || (i==y && j==x)  {
						fmt.Print("o")
					}else if (i==1 && (j!=1 || j!=x)) || (i==y && (j!=x || j!=x)) {
						fmt.Print("-")
					}else if (i!=1 && j==1) || (i!=y && j==x){
						fmt.Print("|")
					}else{
						fmt.Print(" ")
					}
				}
			}
			fmt.Print("\n")
		}
	}
}