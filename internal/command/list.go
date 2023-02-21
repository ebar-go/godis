package command

import "github.com/ebar-go/godis/internal/store"

type List struct {
	storage *store.Store
}

func (list List) LPush(key string, value ...string) error {
	//TODO implement me
	panic("implement me")
}

func (list List) RPush(key string, value ...string) error {
	//TODO implement me
	panic("implement me")
}

func (list List) LPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) RPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) LRange(key string, start, stop int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) Len(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func NewList(storage *store.Store) *List {
	return &List{storage: storage}
}
