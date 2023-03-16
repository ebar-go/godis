package client

import (
	"bufio"
	"fmt"
	"github.com/ebar-go/ego/utils/runtime"
	"github.com/ebar-go/znet/client"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Debug bool   `json:"debug"`
	Host  string `json:"host"`
	Port  int    `json:"port"`
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

	cli.onSuccess()

	go cli.handle(stopCh)

	runtime.WaitClose(stopCh, cli.onClose)
	return nil
}

func (cli *Client) handle(stopCh <-chan struct{}) {
	inputReader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-stopCh:
			return
		default:
			fmt.Printf("\n>")
			input, err := inputReader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					os.Exit(-1)
				}
				fmt.Printf("Error reading:%v", err)
				os.Exit(-1)
			}

			args := strings.Split(input, " ")
			cmd, err := Parse(args)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			if cli.cfg.Debug {
				fmt.Printf("receive command:%v", cmd)
			}

			if err := cmd.Validate(); err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			_, err = cli.conn.Write(cmd.Serialize())
			if err != nil {
				fmt.Printf("send command:%v\n", err)
				continue
			}

			result := make([]byte, 4096)
			n, err := cli.conn.Read(result)
			if err != nil {
				fmt.Printf("read result:%v\n", err)
				continue
			}

			fmt.Printf("%v", string(result[:n]))

		}
	}
}

func (cli *Client) onClose() {
	fmt.Printf("connection closed")
	cli.conn.Close()
}
func (cli *Client) onSuccess() {
	fmt.Printf("Successfully connected to %s\n", cli.cfg.Address())
	fmt.Printf("--------------------------------")
}
