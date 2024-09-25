package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestWatch(t *testing.T) {
   var visitor visitor_id
   time.Sleep(time.Second)
   err := visitor.New()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   watch, err := visitor.watch()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(watch.watch())
}
