package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)>1{
		os.Args = os.Args[1:]
	}else{
		fmt.Println("File name missing")
		return
	}

	if len(os.Args) == 1 {
		file, err := os.Open(os.Args[0])
		if err != nil {
			fmt.Println(err)
		}else{
			data := make([]byte, 14)
			file.Read(data)
				fmt.Println(string(data))
			file.Close()
		}
	}else{
		fmt.Println("Too many arguments")
	}
}