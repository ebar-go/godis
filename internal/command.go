package internal

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
		SortedSet: command.NewSortedSet(storage),
		List:      command.NewList(storage),
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

	// SPop removes and returns one or more random members from the set value store at key.
	SPop(key string, count int) []string

	// SIsMember returns if member is a member of the set stored at key.
	SIsMember(key string, member string) int

	// SMembers returns all the members of the set value stored at key.
	SMembers(key string) []string
}

type SortedSet interface {
	// ZAdd Adds all the specified members with the specified scores to the sorted set stored at key.
	ZAdd(key string, member string, score float64) int

	// Zcard Returns the sorted set cardinality (number of elements) of the sorted set stored at key.
	ZCard(key string) int64
	ZCount(key string, min, max float64) int64

	// ZRange Returns the specified range of elements in the sorted set
	ZRange(key string, start, stop int64) []string
	ZRangeByScore(key string, min, max float64) ([]string, error)

	// ZRem Removes the specified members from the sorted set stored at key. Non existing members are ignored.
	ZRem(key string, member ...string) int

	// ZScore Returns the score of member in the sorted set at key.
	ZScore(key string, member string) (float64, error)
	ZRank(key string, member string) (int64, error)
}

type List interface {
	// LPush insert all the specified values at the head of the list stored at key
	LPush(key string, value ...string) int

	// RPush insert all the specified values at the head of the list stored at key.
	RPush(key string, value ...string) int

	// LPop Removes and returns the first elements of the list stored at key.
	LPop(key string, count int) []string

	// RPop Removes and returns the last elements of the list stored at key.
	RPop(key string, count int) []string

	// LRange Returns the specified elements of the list stored at key.
	LRange(key string, start, stop int64) []string

	// LLen Returns the length of the list stored at key.
	LLen(key string) uint64
}
