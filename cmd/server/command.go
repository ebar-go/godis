package main

import (
	"fmt"
	"github.com/ebar-go/godis/internal"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "godis",
		Short: "godis is a key-value storage",
		Long: `A Fast key-value storage like redis implement by Golang
					Complete documentation is available at https://godis.github.io`,
		Run: func(cmd *cobra.Command, args []string) {
			server := internal.NewServer()
			server.Run()
		},
	}

	prepareChildCommand(rootCmd)

	return rootCmd
}

var childCommand = []*cobra.Command{
	{
		Use:   "version",
		Short: "Print version information",
		Long:  `All software has versions.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v0.1")
		},
	},
}

func prepareChildCommand(rootCmd *cobra.Command) {
	for _, command := range childCommand {
		rootCmd.AddCommand(command)
	}
}
