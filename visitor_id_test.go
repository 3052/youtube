package youtube

import (
   "fmt"
   "reflect"
   "testing"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   visitor_id{},
}

const test_visitor = "CgtpdWhjUkQwRTJsOCiriNy3BjIKCgJVUxIEGgAgMToMCAEgtfbFn7WFwftm"

func TestVisitorId(t *testing.T) {
   var visitor visitor_id
   err := visitor.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", visitor)
}
