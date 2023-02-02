package command

type Hash interface {
	HSet(key string, filed string, value any) error
	HGet(key string, filed string) (value any, err error)
	HExists(key string, filed string) (bool, error)
	HLen(key string) int64
	HDel(key string, field ...string) error
	HKeys(key string) ([]string, error)
	HGetAll(key string) (map[string]any, error)
}
