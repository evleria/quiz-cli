package list

import (
	"bufio"
	"fmt"
	"github.com/cheynewallace/tabby"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/evleria/quiz-cli/pkg/iostreams"
	"github.com/spf13/cobra"
	"text/tabwriter"
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

	if opts.Verbose {
		printCategoriesVerbose(writer, categories.Categories)
	} else {
		printCategories(writer, categories.Categories)
	}

	return nil
}

func printCategories(writer *bufio.Writer, categories map[string][]config.Variation) {
	for name := range categories {
		fmt.Fprintln(writer, name)
	}
}

func printCategoriesVerbose(writer *bufio.Writer, categories map[string][]config.Variation) {
	t := tabby.NewCustom(tabwriter.NewWriter(writer, 0, 0, 2, ' ', 0))
	t.AddHeader("CATEGORY", "VARIATIONS", "QUESTIONS")
	for name, variations := range categories {
		t.AddLine(name, len(variations), countQuestions(variations))
	}
	t.Print()
}

func countQuestions(variations []config.Variation) int {
	sum := 0
	for _, variation := range variations {
		sum += len(variation.Questions)
	}
	return sum
}
