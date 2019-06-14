package piscine

import "fmt"

func countdown(){
	i := 9
	for i > 0 {
		fmt.Print(i)
		i--
	}
	fmt.Print("\n")
}
