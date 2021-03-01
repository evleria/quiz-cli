package printer

import (
	"bufio"
	"fmt"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/evleria/quiz-cli/pkg/runner"
	"io"
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
		if _, err := fmt.Fprintf(p.writer, "%d. %s\n", i + 1, answer.Answer); err != nil {
			return err
		}
	}

	return p.writer.Flush()
}

func (p *Printer) PrintResult(result runner.QuizResult) error {
	if _, err := fmt.Fprintf(p.writer, "%d/%d (%.1f%%)\n", result.Correct, result.Total, float64(result.Correct)/float64(result.Total)*100); err != nil {
		return err
	}

	return p.writer.Flush()
}

