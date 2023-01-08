package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	store := NewStore()
	t.Run("getNotExistKey", func(t *testing.T) {
		assert.Nil(t, store.Get("notExist"))
	})

	t.Run("setNotExistKey", func(t *testing.T) {
		store.Set("someId", "bar")
		obj := store.Get("someId")
		assert.Equal(t, "bar", obj.String())
	})

	t.Run("setExistedKey", func(t *testing.T) {
		store.Set("foo", "bar")
		store.Set("foo", "bar1")
		obj1 := store.Get("foo")
		assert.Equal(t, "bar1", obj1.String())
	})

	t.Run("setInteger", func(t *testing.T) {
		store.Set("age", 1)
		age := store.Get("age")
		assert.Equal(t, "1", age.String())
	})

}

func TestLen(t *testing.T) {
	store := NewStore()
	store.Set("foo", "bar")
	assert.Equal(t, uint64(3), store.Len("foo"))
}

func TestDel(t *testing.T) {
	store := NewStore()
	store.Set("foo", "bar")
	assert.Equal(t, uint64(3), store.Len("foo"))

	store.Del("foo")
	assert.Nil(t, store.Get("foo"))
}
