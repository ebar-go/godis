package client

import (
	"fmt"
	"github.com/ebar-go/godis/constant"
	"github.com/ebar-go/godis/errors"
	"github.com/ebar-go/znet/codec"
	"strings"
)

type Command struct {
	cmd  string
	args []string
}

func (cmd *Command) String() string {
	return fmt.Sprintf("%s %s", cmd.cmd, strings.Join(cmd.args, " "))
}

func (cmd *Command) Serialize() []byte {
	packet := codec.NewPacket(codec.NewJsonCodec())
	packet.Action = constant.ActionCommand
	packet.Seq = 1

	_ = packet.Marshal(map[string]any{"cmd": cmd.cmd, "args": cmd.args})
	p, _ := packet.Pack()

	return p
}

func (cmd *Command) Validate() error {
	argsCount := len(cmd.args)
	switch cmd.cmd {
	case constant.CommandGet:
		if argsCount != 1 {
			return errors.InvalidParams
		}
	case constant.CommandSet:
		if argsCount != 2 {
			return errors.InvalidParams
		}

	}

	return nil
}

func Parse(args []string) (*Command, error) {
	if len(args) <= 1 {
		return nil, errors.InvalidParams
	}

	cmd := &Command{
		cmd:  args[0],
		args: args[1:],
	}

	return cmd, nil
}
