package youtube

import (
	"fmt"
	"testing"
)

func TestVisitorId(t *testing.T) {
	var visitor visitor_id
	err := visitor.New()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(visitor.id())
}
