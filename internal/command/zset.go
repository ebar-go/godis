package command

import "github.com/ebar-go/godis/internal/store"

type SortedSet struct {
	storage *store.Store
}

func (set SortedSet) ZAdd(key string, member string, score float64) int {
	return set.storage.ZAdd(key, member, score)
}

func (set SortedSet) ZCard(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZCount(key string, min, max float64) int64 {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZRange(key string, start, stop int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZRangeByScore(key string, min, max float64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZRem(key string, member string) error {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZScore(key string, member string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (set SortedSet) ZRank(key string, member string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func NewSortedSet(storage *store.Store) *SortedSet {
	return &SortedSet{storage: storage}
}
