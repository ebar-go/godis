package internal

import (
	"github.com/ebar-go/ego/utils/runtime/signal"
	"github.com/ebar-go/godis/internal/store"
	"github.com/ebar-go/znet"
)

type Server struct {
	Command
}

func NewServer() *Server {
	return &Server{
		Command: NewCommand(store.NewStore()),
	}
}

func (s *Server) Run() error {
	instance := znet.New()
	instance.ListenTCP(":3306")
	return instance.Run(signal.SetupSignalHandler())
}
