package cmd

import (
	"fmt"

	"github.com/japelsin/pplx/utils"
	"github.com/kirsle/configdir"
	"github.com/spf13/cobra"
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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure pplx",
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
	ValidArgs: []string{utils.ApiKeyKey, utils.MaxTokensKey, utils.ModelKey, utils.TemperatureKey},
	Args:      cobra.OnlyValidArgs,
}

var configSetApiKeyCmd = &cobra.Command{
	Use:   utils.ApiKeyKey,
	Short: "Set API key",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("API key", args)
		utils.UpdateConfigValue(utils.ApiKeyKey, value)
	},
}

var configSetMaxTokensCmd = &cobra.Command{
	Use:   utils.MaxTokensKey,
	Short: "Set default max tokens",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("Max tokens", args)
		utils.UpdateConfigValue(utils.MaxTokensKey, value)
	},
}

var configSetModelCmd = &cobra.Command{
	Use:       utils.ModelKey,
	Short:     "Set default model",
	Args:      cobra.RangeArgs(0, 1),
	ValidArgs: utils.AvailableModels,
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureSelectValue("Model", args, utils.AvailableModels)
		utils.UpdateConfigValue(utils.ModelKey, value)
	},
}

var configSetTemperatureCmd = &cobra.Command{
	Use:   utils.TemperatureKey,
	Short: "Set default temperature",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("Temperature", args)
		utils.UpdateConfigValue(utils.TemperatureKey, value)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Config subcommands
	configCmd.AddCommand(configPathCmd)
	configCmd.AddCommand(configResetCmd)
	configCmd.AddCommand(configSetCmd)

	// Set subcommands
	configSetCmd.AddCommand(configSetApiKeyCmd)
	configSetCmd.AddCommand(configSetMaxTokensCmd)
	configSetCmd.AddCommand(configSetModelCmd)
	configSetCmd.AddCommand(configSetTemperatureCmd)
}
