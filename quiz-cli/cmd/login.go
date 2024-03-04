package cmd

import (
	"github.com/spf13/cobra"
)

// loginCmd is the root for GET commands
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "commands to login new users",
	Long:  `The root command for GET requests`,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
