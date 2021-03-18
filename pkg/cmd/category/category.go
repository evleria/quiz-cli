package category

import (
	"github.com/evleria/quiz-cli/pkg/cmd/category/list"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/spf13/cobra"
)

func NewCategoryCmd(factory *cmdutils.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "category",
		Short: "shows information of categories",
	}

	cmd.AddCommand(list.NewListCmd(factory))

	return cmd
}
