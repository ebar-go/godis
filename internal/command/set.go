package command

type Set struct{}

func (set Set) SAdd(key string, member string) error {
	//TODO implement me
	panic("implement me")
}

func (set Set) SRem(key string, member string) error {
	//TODO implement me
	panic("implement me")
}

func (set Set) SCard(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func (set Set) SPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (set Set) SIsMember(key string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (set Set) SMembers(key string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
