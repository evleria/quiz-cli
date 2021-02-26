package root

import (
	"github.com/evleria/quiz-cli/pkg/cmd/start"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quiz",
		Short: "challenges your knowledge",
	}

	cmd.AddCommand(start.NewStartCmd())

	return cmd
}
