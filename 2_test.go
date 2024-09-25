package youtube

import (
   "fmt"
   "testing"
   "time"
)

func TestWatch(t *testing.T) {
   time.Sleep(time.Second)
   id, err := visitor_id()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   watch, err := get_watch(id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(watch)
}
