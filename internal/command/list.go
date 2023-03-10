package command

import "github.com/ebar-go/godis/internal/store"

type List struct {
	storage *store.Store
}

func (list List) LPush(key string, value ...string) int {
	return list.storage.LPush(key, value...)
}

func (list List) RPush(key string, value ...string) int {
	return list.storage.RPush(key, value...)
}

func (list List) LPop(key string, count int) []string {
	return list.storage.LPop(key, count)
}

func (list List) RPop(key string, count int) []string {
	return list.storage.RPop(key, count)
}

func (list List) LRange(key string, start, stop int64) []string {
	return list.storage.LRange(key, start, stop)
}

func (list List) LLen(key string) uint64 {
	return list.storage.LLen(key)
}

func NewList(storage *store.Store) *List {
	return &List{storage: storage}
}
