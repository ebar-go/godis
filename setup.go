package godis

import (
	"github.com/ebar-go/godis/internal/command"
	"github.com/ebar-go/godis/internal/store"
	"time"
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

func (s *Server) Exists(key string) (bool, error) {
	return s.key.Exists(key)
}

func (s *Server) Expire(key string, ttl time.Duration) error {
	return s.key.Expire(key, ttl)
}

func (s *Server) TTL(key string) (time.Duration, error) {
	return s.key.TTL(key)
}
