package main

import (
   "fmt"
   "os"
   "slices"
   "strconv"
   "strings"
)

func main() {
   in, err := os.ReadFile("in.txt")
   if err != nil {
      panic(err)
   }
   out, err := os.Create("out.txt")
   if err != nil {
      panic(err)
   }
   defer out.Close()
   var clients []youtube_client
   for _, line := range strings.Fields(string(in)) {
      var client youtube_client
      err = client.Set(line)
      if err != nil {
         panic(err)
      }
      clients = append(clients, client)
   }
   slices.SortFunc(clients, func(a, b youtube_client) int {
      return a.id - b.id
   })
   for _, client := range clients {
      fmt.Fprintln(out, client)
   }
}

func (y *youtube_client) Set(text string) error {
   y.name, text, _ = strings.Cut(text, "(")
   id, _, _ := strings.Cut(text, ")")
   var err error
   y.id, err = strconv.Atoi(id)
   if err != nil {
      return err
   }
   return nil
}

type youtube_client struct {
   name string
   id int
}
