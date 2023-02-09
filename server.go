package godis

import (
	"github.com/ebar-go/godis/internal/store"
)

type Server struct {
	Command
}

func NewServer() *Server {
	return &Server{
		Command: NewCommand(store.NewStore()),
	}
}
