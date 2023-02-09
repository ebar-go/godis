package command

import (
	"github.com/ebar-go/godis/constant"
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/godis/internal/store"
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

func (cmd KeyCommand) Expire(key string, ttl int64) error {
	if !cmd.storage.Has(key) {
		return errors.Nil
	}

	cmd.storage.SetExpire(key, ttl)
	return nil
}

func (cmd KeyCommand) TTL(key string) int64 {
	if !cmd.storage.Has(key) {
		return constant.ExpireResultOfNotFound
	}

	return cmd.storage.GetExpire(key)
}

func (cmd KeyCommand) Exists(key string) bool {
	exist := cmd.storage.Has(key)
	return exist
}
