package main

import (
	"fmt"
	"github.com/ebar-go/godis/cmd/server/options"
	"github.com/ebar-go/godis/internal"
	"github.com/spf13/cobra"
	"log"
)

func NewCommand() *cobra.Command {
	opts := options.NewServerRunOptions()
	rootCmd := &cobra.Command{
		Use:   "godis",
		Short: "godis is a key-value storage",
		Long: `A Fast key-value storage like redis implement by Golang
					Complete documentation is available at https://godis.github.io`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(opts.Address)
			server := internal.NewServer()
			server.Run()
		},
	}

	prepareChildCommand(rootCmd)
	parseFlag(rootCmd, opts)

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

func parseFlag(rootCmd *cobra.Command, opts *options.Options) {
	rootCmd.Flags().StringVar(&opts.Address, "address", ":3306", "set server run address")
}
