package command

import "github.com/ebar-go/godis/internal/store"

type Hash struct {
	storage *store.Store
}

func (hash Hash) HSet(key string, filed string, value any) error {
	return hash.storage.HSet(key, filed, value)
}

func (hash Hash) HGet(key string, filed string) (value any, err error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HExists(key string, filed string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HLen(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HDel(key string, field ...string) error {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HKeys(key string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HGetAll(key string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}

func NewHash(storage *store.Store) *Hash {
	return &Hash{storage: storage}
}
