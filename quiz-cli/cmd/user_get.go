package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getUserCmd = &cobra.Command{
	Use: "user",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("the user identifier must be supplied")
		}
		return nil
	},
	Short: "user {id} -> get user by email",
	Long: `Retrieve a specific user by email.
 This command retrieves a single user specified by the identifier in args.`,
	SilenceUsage: true,
	Run:          getUserRequest,
}

func getUserRequest(cmd *cobra.Command, args []string) {
	resp, err := QuizClient.GetUser(args[0])
	if err != nil {
		_, _ = fmt.Fprint(cmd.ErrOrStderr(), fmt.Sprintf("error getting share with id %s, %s", args[0], err))
		return
	}

	fmt.Printf("email - %s \nscore - %d \ntookquiz - %t \n", resp.Email, resp.Score, resp.TookQuiz)
}

func init() {
	getCmd.AddCommand(getUserCmd)
}
