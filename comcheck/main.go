package main

import (
	"fmt"
	"os"
)

func main() {
	var arrayRune []rune
	var result string
	safeWords := []string{"01", "galaxy", "galaxy 01"}

	for i := 1; i < len(os.Args); i++ {
		arrayRune = []rune(os.Args[i])
	}

	for j := 0; j < len(arrayRune); j++ {
		if arrayRune[j] != ' ' {
			result += string(arrayRune[j])
		}
	}

	for _, s := range safeWords {
		if result == s {
			fmt.Println("Alert!!!")
		}
	}
}