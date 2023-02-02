package command

type String interface {
	Set(key string, value any) error
	Get(key string) (value string, err error)
}
