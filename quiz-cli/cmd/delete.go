package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Command to delete questions or users",
	Long:  `The root command for DELETE requests`,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
