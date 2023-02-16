package godis

import (
	"github.com/ebar-go/godis/internal/command"
	"github.com/ebar-go/godis/internal/store"
)

type Command interface {
	Key
	String
	Hash
	List
	NormalSet
	SortedSet
}

var _ Command = &CommandGroup{}

type CommandGroup struct {
	Key
	String
	Hash
	List
	NormalSet
	SortedSet
}

func NewCommand(storage *store.Store) *CommandGroup {
	return &CommandGroup{
		Key:       command.NewKey(storage),
		String:    command.NewString(storage),
		Hash:      command.NewHash(storage),
		NormalSet: command.NewSet(storage),
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
	// Set sets key to hold the string value.
	// If key already holds a value, it is overwritten, regardless of its type
	Set(key string, value any)

	// Get the value of key. If the key does not exist the special value nil is returned
	Get(key string) (value string)
}

type Hash interface {
	// HSet sets the specified fields to their respective values in the hash stored at key.
	HSet(key string, filed string, value any) error

	// HGet Returns the value associated with field in the hash stored at key.
	HGet(key string, filed string) (value any, err error)

	// HExists returns if field is an existing field in the hash stored at key.
	HExists(key string, filed string) bool

	// HLen Returns the number of fields contained in the hash stored at key.
	HLen(key string) int64

	// HDel Removes the specified fields from the hash stored at key.
	// Specified fields that do not exist within this hash are ignored.
	// If key does not exist, it is treated as an empty hash and this command returns 0.
	HDel(key string, field ...string) int

	// HKeys Returns all field names in the hash stored at key.
	HKeys(key string) []string

	// HGetAll Returns all fields and values of the hash stored at key.
	HGetAll(key string) map[string]any
}

type NormalSet interface {
	// SAdd add the specified members to the set stored at key.
	// If key does not exist, a new set is created before adding the specified members.
	// An error is returned when the value stored at key is not a set.
	SAdd(key string, members ...string) error

	// SRem remove the specified members from the set stored at key.
	// If key does not exist, it is treated as an empty set and this command returns 0.
	// An error is returned when the value stored at key is not a set.
	SRem(key string, members ...string) (int, error)

	// SCard returns the set cardinality (number of elements) of the set stored at key.
	SCard(key string) int64
	SPop(key string) (string, error)
	SIsMember(key string) (bool, error)
	SMembers(key string) ([]string, error)
}

type SortedSet interface {
	ZAdd(key string, member string, score float64) error
	ZCard(key string) int64
	ZCount(key string, min, max float64) int64
	ZRange(key string, start, stop int64) ([]string, error)
	ZRangeByScore(key string, min, max float64) ([]string, error)
	ZRem(key string, member string) error
	ZScore(key string, member string) (float64, error)
	ZRank(key string, member string) (int64, error)
}

type List interface {
	LPush(key string, value ...string) error
	RPush(key string, value ...string) error
	LPop(key string) (string, error)
	RPop(key string) (string, error)
	LRange(key string, start, stop int64) ([]string, error)
	Len(key string) int64
}
