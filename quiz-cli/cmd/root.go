/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"quiz/quiz-cli/client"

	"github.com/spf13/cobra"
)

var (
	QuizClient *client.QuizClient
)

const (
	BaseUrlPath       = "http://localhost:8080"
	QuestionPath      = "/questions/"
	UserPath          = "/users/"
	AdminUserPath     = "/admin/users/"
	AdminQuestionPath = "/admin/questions/"
	LoginPath         = "/login/"
	RegisterPath      = "/register/"
	MethodGet         = "GET"
	MethodPost        = "POST"
	MethodPatch       = "PATCH"
	MethodDelete      = "DELETE"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Quiz servise",
	Long:  `Command line tools for REST requests to Quiz service`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func clientConfig() *client.Config {
	return &client.Config{
		BasePath:     BaseUrlPath,
		QuestionPath: QuestionPath,
		UserPath:     UserPath,
		LoginPath:    LoginPath,
		RegisterPath: RegisterPath,
		MethodGet:    MethodGet,
		MethodDelete: MethodDelete,
		MethodPatch:  MethodPatch,
		MethodPost:   MethodPost,
	}
}

func initClient() {
	if QuizClient == nil {
		var err error
		QuizClient, err = client.NewQuizClient(clientConfig)
		if err != nil {
			fmt.Printf("failed to create client: %s\n", err) // nolint:staticcheck
			os.Exit(1)
		}
	}
}

func init() {
	cobra.OnInitialize(initClient)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.quiz-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// flags.BindCommonFlags(rootCmd)
}
