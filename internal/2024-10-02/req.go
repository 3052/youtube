package main

import (
   "os"
   "fmt"
   "bytes"
)

func main() {
   data, err := os.ReadFile("req.txt")
   if err != nil {
      panic(err)
   }
   for _, line := range bytes.Split(data, []byte("\r\n")) {
      fmt.Printf("%q\n", line)
   }
}
