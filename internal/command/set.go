package command

import "github.com/ebar-go/godis/internal/store"

type Set struct {
	storage *store.Store
}

func (set Set) SAdd(key string, members ...string) error {
	return set.storage.SAdd(key, members...)
}

func (set Set) SRem(key string, members ...string) error {
	//TODO implement me
	panic("implement me")
}

func (set Set) SCard(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func (set Set) SPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (set Set) SIsMember(key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (set Set) SMembers(key string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func NewSet(storage *store.Store) *Set {
	return &Set{storage}
}
