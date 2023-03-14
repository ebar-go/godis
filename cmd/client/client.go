package main

import (
	"fmt"
	"github.com/ebar-go/ego/utils/runtime/signal"
	"github.com/ebar-go/godis/cmd/client/options"
	"github.com/ebar-go/godis/internal/client"
	"github.com/spf13/cobra"
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
			Run(opts)
		},
	}

	prepareChildCommand(rootCmd)
	parseFlag(rootCmd, opts)

	return rootCmd
}

func Run(opts *options.Options) {
	cfg := new(client.Config)
	opts.ApplyTo(cfg)

	cli := client.New(cfg)
	if err := cli.Run(signal.SetupSignalHandler()); err != nil {
		fmt.Printf("error:%v\n", err)
	}
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
