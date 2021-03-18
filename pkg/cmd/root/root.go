package root

import (
	"github.com/evleria/quiz-cli/pkg/cmd/category"
	"github.com/evleria/quiz-cli/pkg/cmd/start"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewRootCmd(factory *cmdutils.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quiz",
		Short: "challenges your knowledge",
	}

	cmd.AddCommand(start.NewStartCmd(factory))
	cmd.AddCommand(category.NewCategoryCmd(factory))

	return cmd
}
