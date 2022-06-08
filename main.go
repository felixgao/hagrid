/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/felixgao/hagrid/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

var configFile string

func initConfig() {

}
