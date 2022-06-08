/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/felixgao/hagrid/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get the secret you stored",
	Long: `Get the secret for the given user and service from the keychain. 
For example:

'<cmd> get [--service myService] [--user yourUser]"'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("get called")
		service := viper.GetString("service")
		user := viper.GetString("user")
		fmt.Printf("cmd flags service: %s, user: %s\n", service, user)
		secret, err := keyring.Get(service, user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", secret)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.
	// Flags

	// @NOTE: Do not set the default values here, that doesn't work correctly!
	getCmd.Flags().StringP("service", "s", "", "service name to be used inside the keychain")
	getCmd.Flags().StringP("user", "u", "", "The user that owns the keychain")

	// Making Flags Required
	// avaible in the future V1.5
	// setCmd.MarkFlagsRequiredTogether("user", "service")
	getCmd.MarkPersistentFlagRequired("user")
	getCmd.MarkPersistentFlagRequired("service")

	//	The magic is setup from root.go where the environment prefix is enforced
	// the Bind Flag name is automatically upper cased
	viper.BindPFlag("service", getCmd.Flags().Lookup("service"))
	viper.BindPFlag("user", getCmd.Flags().Lookup("user"))

	viper.SetDefault("service", "hagrid")
	viper.SetDefault("user", helper.GetCurrentUser())
}
