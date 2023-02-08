package command

import (
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/godis/internal/store"
	"time"
)

type KeyCommand struct {
	storage *store.Store
}

func NewKey(storage *store.Store) *KeyCommand {
	return &KeyCommand{storage: storage}
}

func (cmd KeyCommand) Del(key string) (n uint) {
	return cmd.storage.Del(key)

}

func (cmd KeyCommand) Expire(key string, ttl time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (cmd KeyCommand) Exists(key string) (bool, error) {
	val := cmd.storage.Get(key)
	if val == nil {
		return false, errors.Nil
	}

	return true, nil
}
