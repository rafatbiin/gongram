package main

import (
	"encoding/json"
	"fmt"
	"github.com/rafatbiin/gongram"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewNgramCmd() *cobra.Command {
	var gramSize int
	cmd := &cobra.Command{
		Use:   "ngram",
		Short: "generates Ngrams of given gram size from given text",
		Example: strings.TrimLeft(`Command: gongram ngram -g 4 It’s not about ideas. It’s about making ideas happen!
Response: ["Its not about ideas","not about ideas Its","about ideas Its about","ideas Its about making","Its about making ideas","about making ideas happen"]`, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				text := strings.Join(args, " ")
				if cmd.Flag("gram").Changed {
					generateNgram(text, gramSize)
					return
				}
				generateNgram(text, 1)
			} else {
				fmt.Fprintln(os.Stderr, "please give give some text to generate Ngrams")
				cmd.Help()
			}
		},
	}
	cmd.Flags().IntVarP(&gramSize, "gram", "g", 0, "gram size")
	return cmd
}

func generateNgram(text string, gramSize int ) {
	ngrams, err := gongram.Generate(text, gramSize)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to generate Ngram, error: "+err.Error())
	}
	fngrams, _ := json.Marshal(ngrams)
	fmt.Fprintln(os.Stdout, string(fngrams))
}
