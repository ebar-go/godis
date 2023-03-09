package main

import (
	"fmt"
	"github.com/ebar-go/godis/cmd/client/options"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	cmd := NewCommand()
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewCommand() *cobra.Command {
	opts := options.NewClientOptions()
	rootCmd := &cobra.Command{
		Use:   "godis-cli",
		Short: "godis-cli is a client for godis",
		Long:  `interactive command`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(opts.Host)
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
	rootCmd.Flags().StringVar(&opts.Host, "host", "127.0.0.1", "provide server host")
	rootCmd.Flags().IntVar(&opts.Port, "port", 3306, "provide server port")
}
