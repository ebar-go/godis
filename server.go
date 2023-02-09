package godis

import (
	"github.com/ebar-go/godis/internal/command"
	"github.com/ebar-go/godis/internal/store"
)

type Server struct {
	Key

	String
}

func NewServer() *Server {
	storage := store.NewStore()
	return &Server{
		Key:    command.NewKey(storage),
		String: command.NewString(storage),
	}
}
