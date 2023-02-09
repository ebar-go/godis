package godis

import "time"

type Key interface {
	Del(key string) uint
	Expire(key string, ttl time.Duration) error
	TTL(key string) (time.Duration, error)
	Exists(key string) bool
}
