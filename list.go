package godis

type List interface {
	LPush(key string, value ...string) error
	RPush(key string, value ...string) error
	LPop(key string) (string, error)
	RPop(key string) (string, error)
	LRange(key string, start, stop int64) ([]string, error)
	Len(key string) int64
}
