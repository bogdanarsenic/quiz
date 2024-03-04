package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var getQuestionCmd = &cobra.Command{
	Use: "question",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("the question identifier must be supplied")
		}
		return nil
	},
	Short: "question {id} -> retrieve a specific question by id",
	Long: `Retrieve a specific question by id.
 This command retrieves a single question specified by the identifier in args.`,
	SilenceUsage: true,
	Run:          getQuestionRequest,
}

func getQuestionRequest(cmd *cobra.Command, args []string) {
	resp, err := QuizClient.GetQuestion(args[0])
	if err != nil {
		_, _ = fmt.Fprint(cmd.ErrOrStderr(), fmt.Sprintf("error getting share with id %s, %s", args[0], err))
		return
	}

	fmt.Printf("%d) %s\nAnswers: %v \n", resp.ID, resp.Question, resp.Answers)
}

func init() {
	getCmd.AddCommand(getQuestionCmd)
}
