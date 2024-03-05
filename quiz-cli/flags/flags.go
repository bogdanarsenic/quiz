package flags

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configFlag   = "config"
	usernameFlag = "username"
	passwordFlag = "password"
)

// CommonFlagConfig is a type for global flags
type CommonFlagConfig struct {
	configFile string
	username   string
	password   string
}

// CommonFlags global flags
var CommonFlags = CommonFlagConfig{}

// BindPFlag binds the flag "name" to the command "cmd".
// Binding the flag means that cobra will pick up the value of the flag from the config file if present,
// but override it from the command line if a value for the flag is given there.
func BindPFlag(cmd *cobra.Command, name string) {
	err := viper.BindPFlag(name, cmd.PersistentFlags().Lookup(name))
	if err != nil {
		panic(fmt.Sprintf("failed to initialise config setting %s: %s", name, err))
	}
}

// ConfigFile set config file path
func ConfigFile(file string) {
	CommonFlags.configFile = file
}

// Config returns main config filepath
func Config() string {
	return viper.GetString(configFlag)
}
func Username() string {
	return viper.GetString(usernameFlag)
}

// Name returns name of share or share target
func Password() string {
	return viper.GetString(passwordFlag)
}

// BindCommonFlags bind flags to command
func BindCommonFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&CommonFlags.configFile, configFlag, "", "config file (default is $HOME/.rfos.yaml)")
	cmd.PersistentFlags().StringVarP(&CommonFlags.username, usernameFlag, "u", "user2@gmail.com", "username provided for auth")
	cmd.PersistentFlags().StringVarP(&CommonFlags.password, passwordFlag, "p", "User123!", "password provided for auth")

	BindPFlag(cmd, configFlag)
	BindPFlag(cmd, usernameFlag)
	BindPFlag(cmd, passwordFlag)
}
