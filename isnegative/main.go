package main

import "fmt"

func IsNegative(nombre int) {
	if nombre < 0 {
		fmt.Println("T")
	}else{
		fmt.Println("F")
	}
}
