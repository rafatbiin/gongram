package main

import (
	"github.com/spf13/cobra"
	"os"
)

func main() {
	cmd := NewCmdRoot()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gongram",
		Short: "command line tool for gongram",
		Long: `It's a CLI tool for Go Ngram generator library gongram,
which generates Ngrams of given gram size from given text.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(NewNgramCmd())
	return rootCmd
}
