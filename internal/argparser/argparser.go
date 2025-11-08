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
}

var ParsedOptions *Options

func ParseOptions(rootCmd *cobra.Command) *Options {
	parsedOptions := Options{
		IgnoreCase: false,
		Count:      false,
	}

	rootCmd.Flags().BoolVarP(&parsedOptions.IgnoreCase, "ignore-case", "i", parsedOptions.IgnoreCase, "Ignore case sensitivity in patterns and data")
	rootCmd.Flags().BoolVarP(&parsedOptions.Count, "count", "c", parsedOptions.IgnoreCase, "Print the count of selected lines per file")

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
