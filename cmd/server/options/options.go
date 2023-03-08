package options

type Options struct {
	Address string
}

func NewServerRunOptions() *Options {
	return &Options{
		Address: "",
	}
}
