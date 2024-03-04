package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd is the root for GET commands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "commands to list questions or users",
	Long:  `The root command for GET requests`,
}

func init() {
	rootCmd.AddCommand(listQuestionsCmd)
	rootCmd.AddCommand(listUsersCmd)
}
