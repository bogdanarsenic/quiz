package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "users -> list all users that have taken the test",
	Long: `Retrieve all users except admins.
 This command retrieves all users registered in the service and that has taken the quiz.`,
	SilenceUsage: true,
	Run:          listUserRequest,
}

func listUserRequest(cmd *cobra.Command, args []string) {
	resp, err := QuizClient.ListUsers()
	if err != nil {
		_, _ = fmt.Fprint(cmd.ErrOrStderr(), fmt.Sprintf("error getting users - %s", err))
		return
	}
	fmt.Println("Users that has taken the test: ")
	if len(*resp) < 1 {
		fmt.Println("There are no users that has taken the test!")
	}
	for _, user := range *resp {
		fmt.Printf("%s has the score %d \n", user.Email, user.Score)
	}
}

func init() {
	getCmd.AddCommand(listUsersCmd)
}
