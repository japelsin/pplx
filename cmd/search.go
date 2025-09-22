package cmd

import (
	"fmt"

	"github.com/japelsin/pplx/config"
	"github.com/japelsin/pplx/constants"
	"github.com/japelsin/pplx/perplexity"
	"github.com/japelsin/pplx/validation"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Argument struct {
	key          string
	shorthand    string
	defaultValue any
	description  string
	configKey    string
	auto         bool // Whether to automatically append argument to payload
}

var ARGUMENTS = []Argument{
	{"query", "q", "", "Your query (you'll be prompted if not provided)", "", false},
	{"system_prompt", "s", 1000, "Token limit per request", constants.SYSTEM_PROMPT_CONFIG_KEY, false},
	{"model", "m", "sonar", "Model to use (e.g., sonar, sonar-pro)", constants.MODEL_CONFIG_KEY, true},
	{"max_tokens", "l", 1000, "Token limit per request", constants.MAX_TOKENS_CONFIG_KEY, true},
	{"search_recency_filter", "", "", "Recency filter for search results ('month', 'week', 'day', 'hour')", "", true},
	{"stream", "", true, "Whether to stream response", "", true},
	{"disable_search", "", false, "Disable search", "", true},
	{"enable_search_classifier", "", false, "Enable search classifier", "", true},
	{"frequency_penalty", "", 0.0, "Frequency penalty (0.0 to 1.0)", "", true},
	{"last_updated_after_filter", "", "", "Filter results last updated after this date (YYYY-MM-DD)", "", true},
	{"last_updated_before_filter", "", "", "Filter results last updated before this date (YYYY-MM-DD)", "", true},
	{"presence_penalty", "", 0.0, "Presence penalty (-2.0 to 2.0)", "", true},
	{"return_related_questions", "", false, "Return related questions", "", true},
	{"search_after_date_filter", "", "", "Filter search results after this date (YYYY-MM-DD)", "", true},
	{"search_before_date_filter", "", "", "Filter search results before this date (YYYY-MM-DD)", "", true},
	{"search_domain_filter", "", []string{}, "Filter search results by domain", "", true},
	{"search_mode", "", "", "Search mode ('academic' or 'web')", "", true},
	{"temperature", "", 0.0, "Sampling temperature", "", true},
	{"top_k", "", 0, "Number of search results to consider", "", true},
	{"top_p", "", 1.0, "Nucleus sampling probability", "", true},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search using Perplexity",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
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

		queryFlag := cmd.Flags().Lookup("query")
		if !queryFlag.Changed {
			prompt := promptui.Prompt{
				Label:    "Query",
				Validate: validation.ValidateRequired,
			}

			promptQuery, err := prompt.Run()
			cobra.CheckErr(err)

			cmd.Flags().Set("query", promptQuery)
		}

		streamFlag := cmd.Flags().Lookup("stream")
		if !streamFlag.Changed {
			cmd.Flags().Set("stream", "true")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := viper.GetString(constants.API_KEY_CONFIG_KEY)
		pplxClient := perplexity.NewClient(apiKey)

		systemPrompt := config.GetSystemPrompt()
		if systemPrompt != "" {
			pplxClient.AppendMessage("system", systemPrompt)
		}

		query, err := cmd.Flags().GetString("query")
		cobra.CheckErr(err)
		pplxClient.AppendMessage("user", query)

		pplxClient.SetPayload("model", viper.GetString(constants.MODEL_CONFIG_KEY))
		pplxClient.SetPayload("max_tokens", viper.GetInt(constants.MAX_TOKENS_CONFIG_KEY))

		for _, arg := range ARGUMENTS {
			if !arg.auto {
				continue
			}

			flag := cmd.Flags().Lookup(arg.key)
			if flag != nil && flag.Changed {
				pplxClient.SetPayload(arg.key, flag.Value)
			}
		}

		stream, err := cmd.Flags().GetBool("stream")
		cobra.CheckErr(err)

		var result *perplexity.Result
		var error error

		if stream {
			result, error = pplxClient.MakeStreamedRequest(func(res string) {
				fmt.Print(res)
			})
			cobra.CheckErr(error)
		} else {
			result, error = pplxClient.MakeRequest()
			cobra.CheckErr(error)
			fmt.Println(result.Choices[0].Message.Content)
		}

		// if len(result.SearchResults) > 0 {
		// 	fmt.Println("\n\nSources:")
		// 	for i, source := range result.SearchResults {
		// 		fmt.Printf("[%d] %s - %s\n", i+1, source.Title, source.Url)
		// 	}
		// }

		fmt.Printf("\n\nTotal tokens: %d | Prompt: %d | Completion: %d\n", result.Usage.TotalTokens, result.Usage.PromptTokens, result.Usage.CompletionTokens)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	for _, arg := range ARGUMENTS {
		switch v := arg.defaultValue.(type) {
		case int:
			searchCmd.Flags().IntP(arg.key, arg.shorthand, v, arg.description)
		case float64:
			searchCmd.Flags().Float64P(arg.key, arg.shorthand, v, arg.description)
		case bool:
			searchCmd.Flags().BoolP(arg.key, arg.shorthand, v, arg.description)
		case []string:
			searchCmd.Flags().StringArrayP(arg.key, arg.shorthand, v, arg.description)
		case string:
			searchCmd.Flags().StringP(arg.key, arg.shorthand, v, arg.description)
		}

		if arg.configKey != "" {
			viper.BindPFlag(arg.configKey, searchCmd.Flags().Lookup(arg.key))
		}
	}
}
