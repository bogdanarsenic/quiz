package cmd

import (
	"github.com/spf13/cobra"
)

// loginCmd is the root for GET commands
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "commands to login new users",
	Long:  `command for logging into the service to try quiz`,
	Run:   loginRequest,
}

func loginRequest(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(loginCmd)
}
