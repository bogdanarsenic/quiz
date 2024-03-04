package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd is the root for UPDATE commands
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "commands to update question or user",
	Long:  `The root command for PATCH requests`,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
