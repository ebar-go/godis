package command

import "github.com/ebar-go/godis/internal/store"

type Set struct {
	storage *store.Store
}

func (set Set) SAdd(key string, members ...string) error {
	return set.storage.SAdd(key, members...)
}

func (set Set) SRem(key string, members ...string) (int, error) {
	return set.storage.SRem(key, members...)
}

func (set Set) SCard(key string) int64 {
	return set.storage.SCard(key)
}

func (set Set) SPop(key string, count int) []string {
	return set.storage.SPop(key, count)
}

func (set Set) SIsMember(key string, member string) int {
	return set.storage.SIsMember(key, member)
}

func (set Set) SMembers(key string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func NewSet(storage *store.Store) *Set {
	return &Set{storage}
}
