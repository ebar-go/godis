package types

import (
	"fmt"
	"testing"
)

func TestQuickList(t *testing.T) {
	ql := NewQuickList()
	ql.PushHead(1)
	ql.PushHead(2)
	ql.PushHead(3)

	result := ql.LRange(0, 4)
	for _, entry := range result {
		fmt.Println(entry)
	}
}
