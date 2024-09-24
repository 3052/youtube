package youtube

import (
   "fmt"
   "testing"
)

func TestVisitorId(t *testing.T) {
   visitor, err := visitor_id()
   if err != nil {
      t.Fatal(err)
   }
   visitor, _ = visitor.Get(1)()
   id, _ := visitor.GetBytes(2)()
   fmt.Println(string(id))
}
