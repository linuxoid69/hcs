package helpers

import (
	"os"

	"github.com/spf13/cobra"
)

func DefaultHelp[T []string | []int](cmd *cobra.Command, args *T) {
	if len(*args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
