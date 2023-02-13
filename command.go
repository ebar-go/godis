package godis

import (
	"github.com/ebar-go/godis/internal/command"
	"github.com/ebar-go/godis/internal/store"
)

type Command interface {
	Key
	String
	Hash
}

type CommandGroup struct {
	Key
	String
	Hash
}

func NewCommand(storage *store.Store) Command {
	return &CommandGroup{
		Key:    command.NewKey(storage),
		String: command.NewString(storage),
		Hash:   command.NewHash(storage),
	}
}

type Key interface {
	// Del deletes some keys
	Del(key ...string) uint

	// Expire set expiration for the given key
	Expire(key string, ttl int64) error

	// TTL returns the expiration time of the given key
	TTL(key string) int64

	// Exists returns true if the given key exists
	Exists(key string) bool
}

type String interface {
	Set(key string, value any) error
	Get(key string) (value string, err error)
}

type Hash interface {
	HSet(key string, filed string, value any) error
	HGet(key string, filed string) (value any, err error)
	HExists(key string, filed string) (bool, error)
	HLen(key string) int64
	HDel(key string, field ...string) error
	HKeys(key string) ([]string, error)
	HGetAll(key string) (map[string]any, error)
}
