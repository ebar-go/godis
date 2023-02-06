package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	assert.NotNil(t, ll)
}

func TestLinkedList_Len(t *testing.T) {
	ll := NewLinkedList()
	assert.Equal(t, uint64(0), ll.Len())

	value := "foo"
	ll.AddNodeHead(unsafe.Pointer(&value))
	assert.Equal(t, uint64(1), ll.Len())
}

func TestLinkedList_AddNodeTail(t *testing.T) {
	ll := NewLinkedList()

	items := []string{"a", "b", "c", "d"}
	for _, item := range items {
		elem := item
		ll.AddNodeTail(unsafe.Pointer(&elem))
	}
	assert.Equal(t, uint64(len(items)), ll.Len())
	assert.Equal(t, "d", *(*string)(ll.Last().value))
}
