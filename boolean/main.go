package main

import (
	"os"
	"errors"
	"unicode/utf8"
)
func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}



func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		PrintRune(arrayStr[i])
	}
	PrintRune('\n')
}

func isEven() int {
	lengthOfArg := os.Args[1:]
	if len(lengthOfArg) % 2 == 0{
		return 1
	}
	return 0
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven() == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}