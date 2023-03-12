package options

import "github.com/ebar-go/godis/internal/client"

type Options struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func NewClientOptions() *Options {
	return &Options{}
}

func (opts *Options) ApplyTo(cfg *client.Config) {
	cfg.Host = opts.Host
	cfg.Port = opts.Port
	return
}
