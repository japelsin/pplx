package cmd

import (
	"fmt"

	"github.com/japelsin/pplx/config"
	"github.com/japelsin/pplx/constants"
	"github.com/japelsin/pplx/perplexity"
	"github.com/japelsin/pplx/validation"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var query = ""

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search using Perplexity",
	Args:  cobra.NoArgs,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		recencyFilter := cmd.Flags().Lookup(constants.SEARCH_RECENCY_FILTER_KEY)

		if recencyFilter.Changed {
			err := validation.ValidateRecencyFilter(recencyFilter.Value.String())
			if err != nil {
				return err
			}
		}

		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Ensure API key is set

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

		// Ensure query is set

		if query == "" {
			prompt := promptui.Prompt{
				Label:    "Query",
				Validate: validation.ValidateRequired,
			}

			promptQuery, err := prompt.Run()
			cobra.CheckErr(err)

			query = promptQuery
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString(constants.API_KEY_KEY)
		pplxClient := perplexity.NewClient(apiKey)

		systemPrompt := viper.GetString(constants.SYSTEM_PROMPT_KEY)
		if systemPrompt != "" {
			pplxClient.AppendMessage("system", systemPrompt)
		}

		pplxClient.AppendMessage("user", query)
		pplxClient.SetPayload(constants.STREAM_KEY, true)
		pplxClient.SetPayload(constants.MODEL_KEY, viper.GetString(constants.MODEL_KEY))
		pplxClient.SetPayload(constants.MAX_TOKENS_KEY, viper.GetString(constants.MAX_TOKENS_KEY))

		// Skip validation for now
		cmd.Flags().Visit(func(flag *pflag.Flag) {
			pplxClient.SetPayload(flag.Name, flag.Value)
		})

		result, err := pplxClient.MakeStreamedRequest(func(res string) {
			fmt.Print(res)
		})
		cobra.CheckErr(err)

		fmt.Printf("\n\nTotal tokens: %d | Prompt: %d | Completion: %d\n", result.Usage.TotalTokens, result.Usage.PromptTokens, result.Usage.CompletionTokens)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&query, "query", "q", "", "Your query")

	// Config flags
	searchCmd.Flags().IntP(constants.MAX_TOKENS_KEY, "l", 0, "Token limit per request")
	searchCmd.Flags().StringP(constants.MODEL_KEY, "m", "", "Model to use")
	searchCmd.Flags().StringP(constants.SYSTEM_PROMPT_KEY, "s", "", "System prompt")
	searchCmd.Flags().StringP(constants.API_KEY_KEY, "a", "", "API Key")

	viper.BindPFlag(constants.API_KEY_KEY, searchCmd.Flags().Lookup(constants.API_KEY_KEY))
	viper.BindPFlag(constants.MAX_TOKENS_KEY, searchCmd.Flags().Lookup(constants.MAX_TOKENS_KEY))
	viper.BindPFlag(constants.MODEL_KEY, searchCmd.Flags().Lookup(constants.MODEL_KEY))
	viper.BindPFlag(constants.SYSTEM_PROMPT_KEY, searchCmd.Flags().Lookup(constants.SYSTEM_PROMPT_KEY))
	viper.BindPFlag(constants.STREAM_KEY, searchCmd.Flags().Lookup(constants.STREAM_KEY))

	// Optional flags
	searchCmd.Flags().Float64P(constants.FREQUENCY_PENALTY_KEY, "f", 0, "Token frequency penalty [0, 1.0]")
	searchCmd.Flags().Float64P(constants.PRESENCE_PENALTY_KEY, "p", 0, "Token presence penalty [-2.0, 2.0]")
	searchCmd.Flags().Float64P(constants.TEMPERATURE_KEY, "t", 0, "Response randomness [0, 2.0]")
	searchCmd.Flags().Float64P(constants.TOP_P_KEY, "P", 0, "Probability cutoff for token selection [0, 1.0]")
	searchCmd.Flags().IntP(constants.TOP_K_KEY, "K", 0, "Number of tokens to sample from [0, 2048]")
	searchCmd.Flags().StringArrayP(constants.SEARCH_DOMAIN_FILTER_KEY, "d", []string{}, "Domain filter (e.g. '-d https://x.com -d https://y.com')")
	searchCmd.Flags().StringP(constants.SEARCH_RECENCY_FILTER_KEY, "r", "", "Recency filter (month, week, day or hour)")
}
