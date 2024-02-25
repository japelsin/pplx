package cmd

import (
	"fmt"
	"os"

	"github.com/japelsin/pplx/utils"
	"github.com/kirsle/configdir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "pplx",
	Short: "Simple CLI for interfacing with Perplexity search",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	path := configdir.LocalConfig()

	viper.AddConfigPath(path)
	viper.SetConfigName("pplx")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			utils.ResetConfig()
			viper.SafeWriteConfig()
		} else {
			cobra.CheckErr(err)
		}
	}

	if viper.Get(utils.ApiKeyKey) == "" {
		fmt.Println("API key not set")
		fmt.Println("Enter your Perplexity API key to get started")

		apiKey, err := utils.Prompt("API key")
		cobra.CheckErr(err)

		viper.Set(utils.ApiKeyKey, apiKey)
		viper.WriteConfig()
	}
}

func init() {
	initConfig()
}
