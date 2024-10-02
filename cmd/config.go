package cmd

import (
	"fmt"
	"strconv"

	"github.com/japelsin/pplx/config"
	"github.com/japelsin/pplx/constants"
	"github.com/japelsin/pplx/validation"
	"github.com/kirsle/configdir"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func ensureValue(label string, args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validation.ValidateRequired,
	}

	result, err := prompt.Run()
	cobra.CheckErr(err)

	return result
}

func ensureSelectValue(label string, args []string, items []string) string {
	if len(args) > 0 {
		return args[0]
	}

	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
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
		config.Reset()
	},
}

var configSetCmd = &cobra.Command{
	Use:       "set",
	Short:     "Set config values",
	ValidArgs: constants.CONFIG_KEYS,
	Args:      cobra.OnlyValidArgs,
}

var configSetApiKeyCmd = &cobra.Command{
	Use:   constants.API_KEY_KEY,
	Short: "Set API key",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("API key", args)
		config.SetApiKey(value)

		err := config.Save()
		cobra.CheckErr(err)
	},
}

var configSetMaxTokensCmd = &cobra.Command{
	Use:   constants.MAX_TOKENS_KEY,
	Short: "Set default max tokens",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureValue("Max tokens", args)

		v, _ := strconv.Atoi(value) // Already validated
		config.SetMaxTokens(v)

		err := config.Save()
		cobra.CheckErr(err)
	},
}

var configSetModelCmd = &cobra.Command{
	Use:       constants.MODEL_KEY,
	Short:     "Set default model",
	Args:      cobra.RangeArgs(0, 1),
	ValidArgs: constants.AVAILABLE_MODELS,
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureSelectValue("Model", args, constants.AVAILABLE_MODELS)
		config.SetModel(value)

		err := config.Save()
		cobra.CheckErr(err)
	},
}

var STREAM_OPTIONS = []string{"Yes", "No"}

var configSetStreamCmd = &cobra.Command{
	Use:       constants.STREAM_KEY,
	Short:     "Set whether to stream response",
	Args:      cobra.RangeArgs(0, 1),
	ValidArgs: STREAM_OPTIONS,
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureSelectValue("Stream response", args, STREAM_OPTIONS)
		config.SetStream(value == "Yes")

		err := config.Save()
		cobra.CheckErr(err)
	},
}

var configSetSystemPromptCmd = &cobra.Command{
	Use:       constants.SYSTEM_PROMPT_KEY,
	Short:     "Set default system prompt",
	Args:      cobra.RangeArgs(0, 1),
	ValidArgs: constants.AVAILABLE_MODELS,
	Run: func(cmd *cobra.Command, args []string) {
		value := ensureSelectValue("Model", args, constants.AVAILABLE_MODELS)
		config.SetModel(value)

		err := config.Save()
		cobra.CheckErr(err)
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
	configSetCmd.AddCommand(configSetStreamCmd)
	configSetCmd.AddCommand(configSetSystemPromptCmd)
}
