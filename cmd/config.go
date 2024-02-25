package cmd

import (
	"fmt"

	"github.com/japelsin/pplx/utils"
	"github.com/kirsle/configdir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ensureValue(label string, args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	result, err := utils.Prompt(label)
	cobra.CheckErr(err)

	return result
}

func ensureSelectValue(label string, args []string, items []string) string {
	if len(args) > 0 {
		return args[0]
	}

	_, result, err := utils.PromptSelect(label, items)
	cobra.CheckErr(err)

	return result
}

func updateConfigValue(key string, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()

	cobra.CheckErr(err)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure pplx CLI",
}

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Get configuration file path",
	Run: func(cmd *cobra.Command, args []string) {
		path := configdir.LocalConfig()
		fmt.Println(path)
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset config",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ResetConfig()
	},
}

var configSetCmd = &cobra.Command{
	Use:       "set",
	Short:     "Set config value",
	ValidArgs: []string{utils.AdditionalInstructionsKey, utils.ApiKeyKey, utils.ModelKey},
	Args:      cobra.OnlyValidArgs,
}

var configSetAdditionalInstructionsCmd = &cobra.Command{
	Use:   utils.AdditionalInstructionsKey,
	Short: "Set additional instructions to include at the end of each query",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("Additional instructions", args)
		updateConfigValue(utils.AdditionalInstructionsKey, value)
	},
}

var configSetApiKeyCmd = &cobra.Command{
	Use:   "api_key",
	Short: "Set API key",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("API key", args)
		updateConfigValue(utils.ApiKeyKey, value)
	},
}

var (
	availableModels   = []string{"pplx-7b-chat", "pplx-70b-chat", "pplx-7b-online", "pplx-70b-online", "llama-2-70b-chat", "codellama-34b-instruct", "codellama-70b-instruct", "mistral-7b-instruct", "mixtral-8x7b-instruct"}
	configSetModelCmd = &cobra.Command{
		Use:       utils.ModelKey,
		Short:     "Set model",
		Args:      cobra.RangeArgs(0, 1),
		ValidArgs: availableModels,
		Run: func(cmd *cobra.Command, args []string) {
			value := ensureSelectValue("Model", args, availableModels)
			updateConfigValue(utils.ApiKeyKey, value)
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)

	// Config subcommands
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configResetCmd)
	configCmd.AddCommand(configSetCmd)

	// Set subcommands
	configSetCmd.AddCommand(configSetAdditionalInstructionsCmd)
	configSetCmd.AddCommand(configSetApiKeyCmd)
	configSetCmd.AddCommand(configSetModelCmd)
}