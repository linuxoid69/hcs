package helpers

import (
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

func DefaultHelp[T []string | []int](cmd *cobra.Command, args *T) {
	if len(*args) == 0 {
		if err := cmd.Help(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	}
}

func DeleteElementFromSlice(slice []string, target string) []string {
	return slices.DeleteFunc(slice, func(s string) bool {
		return s == target
	})
}
