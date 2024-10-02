package main

import (
   "fmt"
   "os"
   "regexp"
   "slices"
   "strconv"
   "strings"
)

type youtube_client struct {
   id int
   name string
}

func (y *youtube_client) Set(text string) error {
   y.name, text, _ = strings.Cut(text, "(")
   var err error
   y.id, err = strconv.Atoi(text)
   if err != nil {
      return err
   }
   return nil
}

func main() {
   java, err := os.ReadFile("asrr.java")
   if err != nil {
      panic(err)
   }
   file, err := os.Create("clients.txt")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   var clients []youtube_client
   re := regexp.MustCompile(`\S+\(\d+`)
   for _, match := range re.FindAllString(string(java), -1) {
      var client youtube_client
      err = client.Set(match)
      if err != nil {
         panic(err)
      }
      clients = append(clients, client)
   }
   slices.SortFunc(clients, func(a, b youtube_client) int {
      return a.id - b.id
   })
   for _, client := range clients {
      fmt.Fprintln(file, client)
   }
}
