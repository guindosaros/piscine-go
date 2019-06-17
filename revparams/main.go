package main

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args[1:]
	for i:= len(args)-1;i>=0;i--{
		fmt.Println(args[i])
	}
}
