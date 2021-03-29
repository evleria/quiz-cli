package root

import (
	"github.com/evleria/quiz-cli/pkg/cmd/category"
	"github.com/evleria/quiz-cli/pkg/cmd/start"
	"github.com/evleria/quiz-cli/pkg/cmdutils"
	"github.com/evleria/quiz-cli/pkg/config"
	"github.com/spf13/cobra"
)

func NewRootCmd(factory *cmdutils.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quiz",
		Short: "challenges your knowledge",
	}

	var configPath string
	cmd.PersistentFlags().StringVarP(&configPath, "path", "p", config.DefaultPath, "path to config. Remote config starts with \"http://\"")
	factory.ConfigFunc = func() config.Config {
		return config.CreateConfigFromPath(configPath)
	}

	cmd.AddCommand(start.NewStartCmd(factory))
	cmd.AddCommand(category.NewCategoryCmd(factory))

	return cmd
}
