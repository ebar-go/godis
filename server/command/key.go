package command

import "time"

type Key interface {
	Del(key string) error
	Expire(key string, ttl time.Duration) error
	Exists(key string) (bool, error)
}
