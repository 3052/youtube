package youtube

import (
   "154.pages.dev/protobuf"
   "bytes"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "time"
)

func GetWatch() {
   message, _ = message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   message, _ = message.Get(3)()
   address := func() string {
      b, _ := message.GetBytes(2)()
      return string(b)
   }()
   fmt.Println(address)
   after := time.After(30 * time.Second)
   for {
      select {
      case <-time.After(time.Second):
         resp, err := http.Head(address)
         if err != nil {
            panic(err)
         }
         fmt.Println(resp.Status)
      case <-after:
         return
      }
   }
}
