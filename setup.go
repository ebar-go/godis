package godis

import (
	"github.com/ebar-go/godis/internal/command"
	"github.com/ebar-go/godis/internal/store"
)

type Server struct {
	key Key
}

func NewServer() *Server {
	storage := store.NewStore()
	return &Server{
		key: command.NewKey(storage),
	}
}

func (s *Server) Del(key string) uint {
	return s.key.Del(key)
}
