package godis

type Key interface {
	// Del deletes some keys
	Del(key string) uint

	// Expire set expiration for the given key
	Expire(key string, ttl int64) error

	// TTL returns the expiration time of the given key
	TTL(key string) int64

	// Exists returns true if the given key exists
	Exists(key string) bool
}
