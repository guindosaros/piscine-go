package piscine

import "fmt"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := 1; j <= 99; j++ {
			if i<j {
				if i==98 && j==99  {
					fmt.Printf("%02d %02d\n",i,j)
				}else{
					fmt.Printf("%02d %02d, ",i,j)
				}
			}
		}
	}
}