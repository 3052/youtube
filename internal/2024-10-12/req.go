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
   file, err := os.Create("quote.txt")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   for _, line := range bytes.Split(data, []byte("\r\n")) {
      fmt.Fprintf(file, "%q\n", line)
   }
}
