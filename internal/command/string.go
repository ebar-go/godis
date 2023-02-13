package command

import (
	"github.com/ebar-go/godis/internal/store"
)

type String struct {
	storage *store.Store
}

func (s String) Set(key string, value any) {
	s.storage.Set(key, value)
}

func (s String) Get(key string) (value string) {
	obj := s.storage.Get(key)
	if obj == nil {
		return "Nil"
	}

	value = obj.String()
	return
}

func NewString(storage *store.Store) *String {
	return &String{storage: storage}
}
