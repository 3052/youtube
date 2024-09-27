package youtube

import (
   "154.pages.dev/protobuf"
   "fmt"
   "testing"
)

var test_visitor = protobuf.Message{
   1: {protobuf.Message{
      2: {protobuf.Bytes("CgtpdWhjUkQwRTJsOCiriNy3BjIKCgJVUxIEGgAgMToMCAEgtfbFn7WFwftm")},
   }},
}

func TestVisitor(t *testing.T) {
   var visitor visitor_id
   err := visitor.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", visitor.id())
}
