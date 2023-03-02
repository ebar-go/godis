package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := NewCommand()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
