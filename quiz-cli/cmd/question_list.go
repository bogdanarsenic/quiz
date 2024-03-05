package cmd

import (
	"fmt"
	"os"
	models "quiz/quiz-cli/structs"

	"github.com/manifoldco/promptui"
	prompt "github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var listQuestionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "questions -> list of all questions",
	Long: `Retrieve all questions.
 This command retrieves all questions in the quiz.`,
	SilenceUsage: true,
	Run:          listQuestionRequest,
}

func promptGetSelect(pc models.PromptContent, question models.Question, score *int) {

	items := []string{}

	index := -1
	var result string
	var err error

	templates := &prompt.SelectTemplates{
		Label:    "{{ . }}?",
		Inactive: " {{ .Answers | cyan }} ({{ .Answers | red }})",
		Selected: "\U0001F336 {{ .Answers | red | cyan }}",
	}

	for index < 0 {
		prompt := promptui.Select{
			Label:     pc.Label,
			Items:     question.Answers,
			Templates: templates,
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		if result == question.Answer {
			*score++
		}
	}
}

func listQuestionRequest(cmd *cobra.Command, args []string) {

	resp, err := QuizClient.ListQuestions()
	if err != nil || *resp == nil {
		_, _ = fmt.Fprint(cmd.ErrOrStderr(), fmt.Sprintf("Unauthorized error!"))
		return
	}
	score := 0

	for _, q := range *resp {
		catPromptContent := models.PromptContent{
			ErrorMsg: "Please answer the question",
			Label:    fmt.Sprintf("%d) %s", q.ID, q.Question),
		}

		promptGetSelect(catPromptContent, q, &score)
	}
	fmt.Printf("You answer correctly on %d out from %d questions \n", score, len(*resp))

	users, err := QuizClient.ListUsers()
	if err != nil {
		_, _ = fmt.Fprint(cmd.ErrOrStderr(), fmt.Sprintf("error getting users - %s", err))
		return
	}

	if len(*users) < 1 {
		fmt.Println("Better then 100%% of users!")
		return
	}

	numbersOfUsers := 0.0
	peopleWhoTookQuiz := 0.0
	for _, u := range *users {
		if u.TookQuiz {
			peopleWhoTookQuiz++
			if score > u.Score {
				numbersOfUsers++
			}
		}
	}

	if peopleWhoTookQuiz < 1 {
		fmt.Println("Better then 100%% of users!")
		return
	}

	percent := 100.0 * numbersOfUsers / peopleWhoTookQuiz

	fmt.Printf("You are better then %.2f%% of users\n", percent)
}

func init() {
	getCmd.AddCommand(listQuestionsCmd)
}
