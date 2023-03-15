package client

import (
	"fmt"
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

func Parse(args []string) (*Command, error) {
	if len(args) <= 1 {
		return nil, errors.New("invalid params")
	}

	cmd := &Command{
		cmd:  args[0],
		args: args[1:],
	}

	return cmd, nil
}
