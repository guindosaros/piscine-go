package main

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args[1:]
	for _,res:= range args{
		fmt.Println(res)
	}
}
