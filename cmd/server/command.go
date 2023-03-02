package main

import (
	"github.com/ebar-go/godis/internal"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "godis",
		Short: "godis is a key-value storage",
		Long: `A Fast key-value storage like redis implement by Golang
					Complete documentation is available at https://godis.github.io`,
		Run: func(cmd *cobra.Command, args []string) {
			server := internal.NewServer()
			server.Run()
		},
	}
}
