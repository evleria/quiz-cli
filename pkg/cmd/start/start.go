package start

import (
	"bufio"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/evleria/quiz-cli/pkg/iostreams"
	"github.com/evleria/quiz-cli/pkg/quiz"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"strings"
	"unicode"
)

type StartCmdOptions struct {
	Categories []string

	ConfigFunc func() config.Config
	IOStreams  iostreams.IOStreams
}

func NewStartCmd(factory *cmdutils.Factory) *cobra.Command {
	opts := &StartCmdOptions{
		ConfigFunc: factory.ConfigFunc,
		IOStreams:  factory.IOStreams,
	}

	cmd := &cobra.Command{
		Use:   "start",
		Short: "starts a quiz",
		RunE: func(c *cobra.Command, args []string) error {
			return runStart(opts)
		},
	}

	cmd.Flags().StringArrayVarP(&opts.Categories, "category", "c", nil, "filters categories for quiz")

	return cmd
}

func runStart(opts *StartCmdOptions) error {
	categories, err := opts.ConfigFunc().ReadConfig()
	cmdutils.CheckError(err)

	variations := getVariations(categories, opts.Categories)
	questions := pickQuestions(variations)

	scanner := bufio.NewScanner(opts.IOStreams.In)
	printer := quiz.NewPrinter(opts.IOStreams.Out)
	runner := quiz.NewRunner(questions)
	for runner.Next() {
		question := runner.GenerateQuestion()

		if err := cmdutils.ClearConsole(opts.IOStreams.Out); err != nil {
			return err
		}
		if err := printer.PrintQuestion(question); err != nil {
			return err
		}

		if scanner.Scan() {
			runner.MarkAnswer(getAnswers(scanner.Text()))
		} else {
			return scanner.Err()
		}
	}

	if err := cmdutils.ClearConsole(opts.IOStreams.Out); err != nil {
		return err
	}
	if err := printer.PrintResult(runner.Result()); err != nil {
		return err
	}

	return nil
}

func getVariations(categories *config.Categories, filteredCategories []string) []config.Variation {
	var result []config.Variation
	if len(filteredCategories) != 0 {
		for _, filteredCategory := range filteredCategories {
			if variations, ok := categories.Categories[filteredCategory]; ok {
				result = append(result, variations...)
			}
		}
	} else {
		for _, variations := range categories.Categories {
			result = append(result, variations...)
		}
	}

	return result
}

func pickQuestions(variations []config.Variation) []config.Question {
	var result []config.Question

	for _, variation := range variations {
		if l := len(variation.Questions); l > 0 {
			question := variation.Questions[rand.Intn(l)]
			question.Links = append(variation.Links, question.Links...)
			result = append(result, question)
		}
	}

	return result
}

func getAnswers(text string) []int {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsDigit(r)
	})

	var result []int
	for _, word := range words {
		if n, err := strconv.Atoi(word); err == nil {
			result = append(result, n-1)
		}
	}

	return result
}
