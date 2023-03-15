package client

import (
	"fmt"
	"github.com/ebar-go/godis/constant"
	"github.com/ebar-go/godis/errors"
	"strings"
)

type Command struct {
	cmd  string
	args []string
}

func (cmd *Command) String() string {
	return fmt.Sprintf("%s %s", cmd.cmd, strings.Join(cmd.args, " "))
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
