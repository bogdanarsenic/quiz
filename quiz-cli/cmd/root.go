/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"quiz/quiz-cli/client"
	"quiz/quiz-cli/flags"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	QuizClient *client.QuizClient
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Quiz servise",
	Long:  `Command line tools for REST requests to Quiz service`,
}

type rootFlagsConfig struct {
	rootToken string
}

var RootFlags = rootFlagsConfig{}

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
		Token:        viper.GetString(tokenFlag),
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

	rootCmd.PersistentFlags().StringVarP(&RootFlags.rootToken, tokenFlag, "t", "", "bearer token for auth")
	_ = rootCmd.MarkPersistentFlagRequired("token")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.quiz-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	flags.BindPFlag(rootCmd, tokenFlag)
	flags.BindCommonFlags(rootCmd)
}
