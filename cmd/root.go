/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"os"

	"github.com/spf13/cobra"
	
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "github-readme-cli",
	Short: "A brief description of your application",
	Long: `github-readme-cli is a CLI tool that will ask for a repo and return the readme file.`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: false,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
