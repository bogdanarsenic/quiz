package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new question or user",
	Long:  `The root command for POST requests.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
