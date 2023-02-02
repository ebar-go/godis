package command

type Set interface {
	SAdd(key string, member string) error
	SRem(key string, member string) error
	SCard(key string) int
	SPop(key string) (string, error)
	SIsMember(key string) (bool, error)
	SMembers(key string) ([]string, error)
}
