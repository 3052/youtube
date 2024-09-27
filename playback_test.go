package youtube

import (
   "fmt"
   "testing"
)

func TestPlayback(t *testing.T) {
   watch, err := visitor_id{test_visitor}.watch(test_video, nil)
   if err != nil {
      t.Fatal(err)
   }
   next := watch.play()
   for {
      play, ok := next()
      if !ok {
         break
      }
      fmt.Print(play, "\n\n")
   }
}
