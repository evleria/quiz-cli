package quiz

import (
	"bufio"
	"fmt"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/fatih/color"
	"io"
)

var (
	correctColor   = color.New(color.FgGreen).FprintfFunc()
	incorrectColor = color.New(color.FgRed).FprintfFunc()
	noColor        = func(w io.Writer, format string, a ...interface{}) {
		fmt.Fprintf(w, format, a...)
	}
)

type Printer struct {
	writer *bufio.Writer
}

func NewPrinter(writer io.Writer) *Printer {
	return &Printer{
		writer: bufio.NewWriter(writer),
	}
}

func (p *Printer) PrintQuestion(question config.Question) error {
	if _, err := fmt.Fprintf(p.writer, "%s\n", question.Question); err != nil {
		return err
	}
	for i, answer := range question.Answers {
		if _, err := fmt.Fprintf(p.writer, "%d. %s\n", i+1, answer.Answer); err != nil {
			return err
		}
	}

	return p.writer.Flush()
}

func (p *Printer) PrintResult(result Result) error {
	fmt.Fprintf(p.writer, "Stats: %d/%d (%.1f%%)\n", result.Correct, result.Total, float64(result.Correct)/float64(result.Total)*100)
	for _, wq := range result.WrongQuestions {
		fmt.Fprintln(p.writer, wq.Question.Question)
		answered := make(map[int]bool, len(wq.Answer.Indices))
		for _, i := range wq.Answer.Indices {
			answered[i] = true
		}

		for i, answer := range wq.Question.Answers {
			fprintfFunc := noColor
			if answer.IsCorrect {
				fprintfFunc = correctColor
			} else if _, ok := answered[i]; ok {
				fprintfFunc = incorrectColor
			}
			placeholder := getPlaceholder(answer.IsCorrect, answered[i])
			fprintfFunc(p.writer, "%d) %c %s\n", i+1, placeholder, answer.Answer)
		}

		if len(wq.Question.Links) > 0 {
			fmt.Fprintln(p.writer, "More about the topic:")
			for _, link := range wq.Question.Links {
				fmt.Fprintln(p.writer, link)
			}
		}

		fmt.Fprintln(p.writer)
	}

	return p.writer.Flush()
}

func getPlaceholder(correct, picked bool) rune {
	if picked != correct {
		if correct {
			return '+'
		} else {
			return '-'
		}
	} else {
		return ' '
	}
}
