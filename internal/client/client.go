package client

import (
	"fmt"
	"github.com/ebar-go/ego/utils/runtime"
	"github.com/ebar-go/znet/client"
	"net"
	"strconv"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (c *Config) Address() string {
	return net.JoinHostPort(c.Host, strconv.Itoa(c.Port))
}

type Client struct {
	cfg  *Config
	conn net.Conn
}

func New(cfg *Config) *Client {
	return &Client{cfg: cfg}
}

func (cli *Client) Run(stopCh <-chan struct{}) error {
	conn, err := client.DialTCP(cli.cfg.Address())
	if err != nil {
		return err
	}

	cli.conn = conn

	runtime.WaitClose(stopCh, cli.onClose)
	return nil
}

func (cli *Client) onClose() {
	fmt.Printf("connection closed")
	cli.conn.Close()
}
func (cli *Client) onSuccess() {
	fmt.Printf(`Successfully connected to %s\n`, cli.cfg.Address())
}
