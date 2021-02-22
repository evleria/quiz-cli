package start

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "starts a quiz",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("quiz started")
		},
	}

	return cmd
}
