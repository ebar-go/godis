package command

import (
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/godis/internal/store"
)

type String struct {
	storage *store.Store
}

func (s String) Set(key string, value any) error {
	s.storage.Set(key, value)
	return nil
}

func (s String) Get(key string) (value string, err error) {
	obj := s.storage.Get(key)
	if obj == nil {
		err = errors.Nil
		return
	}

	value = obj.String()
	return
}

func NewString(storage *store.Store) *String {
	return &String{storage: storage}
}
