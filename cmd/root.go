package cmd

import (
	"fmt"
	"os"

	"github.com/japelsin/pplx/config"
	"github.com/japelsin/pplx/validation"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pplx",
	Short: "Simple CLI for interfacing with Perplexity's API",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := config.Init()
	cobra.CheckErr(err)

	if config.GetApiKey() == "" {
		fmt.Println("Enter your Perplexity API key to get started")

		prompt := promptui.Prompt{
			Label:    "API key",
			Validate: validation.ValidateRequired,
		}

		apiKey, err := prompt.Run()
		cobra.CheckErr(err)

		config.SetApiKey(apiKey)
		config.Save()
	}

	cobra.CheckErr(err)
}
