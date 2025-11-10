package argparser

import (
	"github.com/spf13/cobra"
)

type Arguments struct {
	Files   []string
	Pattern string
}

type Options struct {
	IgnoreCase bool
	Count      bool
	WholeWord  bool
}

var ParsedOptions *Options

var DefaultOptions = Options{
	IgnoreCase: false,
	Count:      false,
	WholeWord:  false,
}

func ParseOptions(rootCmd *cobra.Command) *Options {
	parsedOptions := DefaultOptions

	rootCmd.Flags().BoolVarP(&parsedOptions.IgnoreCase, "ignore-case", "i", DefaultOptions.IgnoreCase, "Ignore case sensitivity in patterns and data")
	rootCmd.Flags().BoolVarP(&parsedOptions.Count, "count", "c", DefaultOptions.IgnoreCase, "Print the count of selected lines per file")
	rootCmd.Flags().BoolVarP(&parsedOptions.WholeWord, "word-regexp", "w", DefaultOptions.WholeWord, "Match the whole word when pattern matches part of a word")

	return &parsedOptions
}

func ParseArguments(args []string) Arguments {
	parsedArguments := Arguments{}

	if len(args) > 0 {
		parsedArguments.Pattern = args[0]

		if len(args) > 1 {
			parsedArguments.Files = args[1:]
		}
	}

	return parsedArguments
}
