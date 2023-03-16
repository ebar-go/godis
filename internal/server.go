package internal

import (
	"github.com/ebar-go/ego/utils/runtime/signal"
	"github.com/ebar-go/godis/constant"
	"github.com/ebar-go/godis/internal/store"
	"github.com/ebar-go/znet"
	"log"
)

type Server struct {
	Command
}

func NewServer() *Server {
	return &Server{
		Command: NewCommand(store.NewStore()),
	}
}

type Response struct {
	Code int `json:"code"`
}

func (s *Server) Run() error {
	instance := znet.New(func(options *znet.Options) {

	})
	instance.Router().Route(constant.ActionCommand, func(ctx *znet.Context) (any, error) {
		log.Println("request:", string(ctx.Packet().Body))
		return Response{Code: 0}, nil
	})
	instance.ListenTCP(":3306")
	return instance.Run(signal.SetupSignalHandler())
}
