package list

import (
	"bufio"
	"fmt"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/evleria/quiz-cli/pkg/iostreams"
	"github.com/spf13/cobra"
)

type ListCmdOptions struct {
	Verbose bool

	Config    *config.Config
	IOStreams iostreams.IOStreams
}

func NewListCmd(factory *cmdutils.Factory) *cobra.Command {
	opts := &ListCmdOptions{
		Config:    factory.Config,
		IOStreams: factory.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "shows list of categories",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "verbose output")

	return cmd
}

func runList(opts *ListCmdOptions) error {
	categories, err := opts.Config.ReadConfig()
	cmdutils.CheckError(err)

	writer := bufio.NewWriter(opts.IOStreams.Out)
	defer writer.Flush()

	for name, variations := range categories.Categories {
		if opts.Verbose {
			fmt.Fprintf(writer, "%8s | variations: %2d, questions: %2d\n", name, len(variations), countQuestions(variations))
		} else {
			fmt.Fprintln(writer, name)
		}
	}

	return nil
}

func countQuestions(variations []config.Variation) int {
	sum := 0
	for _, variation := range variations {
		sum += len(variation.Questions)
	}
	return sum
}
