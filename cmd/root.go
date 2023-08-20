/*
Copyright © 2023 Rustam Tagaev linuxoid69@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hcs",
	Short: "HCS is a switch of environments.",
	Long:  `HCS is an environment switching program for Hashicorp products.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
