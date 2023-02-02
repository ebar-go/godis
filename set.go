package godis

type Set interface {
	SAdd(key string, member string) error
	SRem(key string, member string) error
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
