package youtube

import (
   "fmt"
   "testing"
)

const test_visitor = "CgtpdWhjUkQwRTJsOCiriNy3BjIKCgJVUxIEGgAgMToMCAEgtfbFn7WFwftm"

func TestVisitorId(t *testing.T) {
   var visitor visitor_id
   err := visitor.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", visitor)
}
