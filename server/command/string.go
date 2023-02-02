package command

import "time"

type String interface {
	Set(key string, value any) error
	Get(key string) (value string, err error)
}

type Delete interface {
	Del(key string) error
}

type Expire interface {
	Expire(key string, ttl time.Duration) error
}
