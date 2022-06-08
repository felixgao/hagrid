/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/felixgao/hagrid/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "set the password to the keychain",
	Long: `Set a password or secret text that can be retrived from the keychain. 

	'<cmd> set --service myService --user yourUser --password "your secret password"'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("set called")
		service := viper.GetString("service")
		user := viper.GetString("user")
		password := viper.GetString("password")
		fmt.Printf("cmd flags service: %s, user: %s, password: %s\n", service, user, password)
		// set password
		err := keyring.Set(service, user, password)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("Successfully wrote your secret to user: %s, service: %s\n", user, service)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.
	// Flags

	// @NOTE: Do not set the default values here, that doesn't work correctly!
	setCmd.Flags().StringP("service", "s", "", "service name to be used inside the keychain")
	setCmd.Flags().StringP("user", "u", "", "The user that owns the keychain")
	setCmd.Flags().StringP("password", "p", "", "Your secret that you want to store to the keychain")

	// Making Flags Required
	// avaible in the future V1.5
	// setCmd.MarkFlagsRequiredTogether("user", "password")
	setCmd.MarkPersistentFlagRequired("user")
	setCmd.MarkPersistentFlagRequired("password")

	//	The magic is setup from root.go where the environment prefix is enforced
	// the Bind Flag name is automatically upper cased
	viper.BindPFlag("service", setCmd.Flags().Lookup("service"))
	viper.BindPFlag("user", setCmd.Flags().Lookup("user"))
	viper.BindPFlag("password", setCmd.Flags().Lookup("password"))

	viper.SetDefault("service", "hagrid")
	viper.SetDefault("user", helper.GetCurrentUser())
}
