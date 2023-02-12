package command

type Hash struct{}

func (hash Hash) HSet(key string, filed string, value any) error {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HGet(key string, filed string) (value any, err error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HExists(key string, filed string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HLen(key string) int64 {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HDel(key string, field ...string) error {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HKeys(key string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (hash Hash) HGetAll(key string) (map[string]any, error) {
	//TODO implement me
	panic("implement me")
}
