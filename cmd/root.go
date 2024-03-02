package cmd

import (
	"os"

	"github.com/japelsin/pplx/utils"
	"github.com/spf13/cobra"
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

func init() {
	utils.InitConfig()
}
