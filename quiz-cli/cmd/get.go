package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd is the root for GET commands
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "commands to get questions or users",
	Long:  `The root command for GET requests`,
}

func init() {
	rootCmd.AddCommand(getQuestionCmd)
	rootCmd.AddCommand(getUserCmd)
}
