package command

type List struct {
}

func (list List) LPush(key string, value ...string) error {
	//TODO implement me
	panic("implement me")
}

func (list List) RPush(key string, value ...string) error {
	//TODO implement me
	panic("implement me")
}

func (list List) LPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) RPop(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) LRange(key string, start, stop int64) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (list List) Len(key string) int64 {
	//TODO implement me
	panic("implement me")
}