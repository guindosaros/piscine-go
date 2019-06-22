
package main

import (
	"fmt"
	"os"
)

func main() {
  if len(os.Args)==1 {
    return
  }else{
    os.Args = os.Args[1:]
    
    for _, v := range os.Args {
        file, err := os.Open(v)
        if err != nil {
          fmt.Println(err)
          return
        }else{
          data := make([]byte,443)
          file.Read(data)
            if len(os.Args) == 1 {
              fmt.Println(string(data))
            }else{
              fmt.Println(string(data), "\n")
            }
          file.Close()
        }
    }
  }
}