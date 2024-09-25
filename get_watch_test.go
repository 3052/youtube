package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestGetWatch(t *testing.T) {
   watch, err := visitor_id{}.watch()
   if err != nil {
      t.Fatal(err)
   }
   message, _ := watch.message.Get(1)()
   message, _ = message.Get(2)()
   message, _ = message.Get(4)()
   message, _ = message.Get(3)()
   address, _ := message.GetBytes(2)()
   fmt.Printf("%q\n", address)
}

func TestWatch(t *testing.T) {
   var visitor visitor_id
   err := visitor.New()
   if err != nil {
      t.Fatal(err)
   }
   watch, err := visitor.watch()
   if err != nil {
      t.Fatal(err)
   }
   message, _ := watch.message.Get(1)()
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
