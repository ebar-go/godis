package options

type Options struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewClientOptions() *Options {
	return &Options{}
}
