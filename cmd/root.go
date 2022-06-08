/*
Copyright © 2022 Felix Gao <>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hagrid",
	Short: "Hagrid is an application manages your eiamCli password",
	Long: `Hagrid is designed to help you manage your eiamCli password in the Mac OS Keychain. 
	For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string
var Service string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hagrid.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readConfig() {
	home, err := homedir.Dir()
	// Maybe use $HOME/.hagrid instead?
	fmt.Printf("Looking for config file at %s\n", home)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	// viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".hagrid")

}

func initConfig() {

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory and set the config to default name of the config
		readConfig()
	}
	viper.SetEnvPrefix("HAGRID")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
	viper.WatchConfig()

	// Search config in home directory with name ".hagrid" (without extension).

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

}
