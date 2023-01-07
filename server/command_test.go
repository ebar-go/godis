package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	store := NewStore()
	assert.Nil(t, store.Get("notExist"))
	
	store.Set("foo", "bar")
	obj := store.Get("foo")
	assert.Equal(t, "bar", obj.String())

	store.Set("foo", "bar1")
	obj1 := store.Get("foo")
	assert.Equal(t, "bar1", obj1.String())

}
