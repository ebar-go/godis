package server

import (
	"log"
	"testing"
)

func TestSet(t *testing.T) {
	store := NewStore()
	store.Set("foo", "bar")
	obj := store.Get("foo")
	log.Println(obj)
}
