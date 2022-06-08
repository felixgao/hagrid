package helper

import (
	"fmt"
	"log"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func HandleError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func LoadConfig() {
	viper.SetEnvPrefix("HAGRID")
	viper.AutomaticEnv()

}

func GetCurrentUser() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return currentUser.Username
}

func ViperCommonBinding(c *cobra.Command) {
	viper.BindPFlag("service", c.Flags().Lookup("service"))
	viper.BindPFlag("user", c.Flags().Lookup("user"))

	viper.SetDefault("service", "hagrid")
	viper.SetDefault("user", GetCurrentUser())
}
