package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/japelsin/pplx/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type perplexityPayloadMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type perplexityPayload map[string]interface{}

type perplexityResult struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created int    `json:"created"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Object  string `json:"object"`
	Choices []struct {
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Delta struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search using Perplexity",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// ****************************************************
		// Parse argument & flags
		// ****************************************************

		query, err := utils.Prompt("Query")
		cobra.CheckErr(err)

		payload := perplexityPayload{}

		payload["stream"] = true
		payload["model"] = viper.Get(utils.ModelKey)

		additionalInstructions := viper.Get(utils.AdditionalInstructionsKey)
		payload["messages"] = []perplexityPayloadMessage{{"user", fmt.Sprintf("%s %s", query, additionalInstructions)}}

		// Skip validation for now
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if flag.Changed {
				payload[flag.Name] = flag.Value.String()
			}
		})

		// ****************************************************
		// Send request
		// ****************************************************

		url := "https://api.perplexity.ai/chat/completions"

		data, err := json.Marshal(payload)
		cobra.CheckErr(err)

		req, err := http.NewRequest("POST", url, bytes.NewReader(data))
		cobra.CheckErr(err)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.Get("api_key")))
		req.Header.Add("Content-Type", "application/json")

		client := http.Client{}

		response, err := client.Do(req)
		cobra.CheckErr(err)

		if response.StatusCode != 200 {
			message := fmt.Sprintf("Request error:  %s", response.Status)
			cobra.CheckErr(errors.New(message))
		}

		// ****************************************************
		// Parse & output result
		// ****************************************************

		result := perplexityResult{}
		scanner := bufio.NewScanner(response.Body)
		prevLen := 0

		for scanner.Scan() {
			bytes := scanner.Bytes()
			bytesLen := len(bytes)

			if bytesLen > 6 {
				err := json.Unmarshal(bytes[6:], &result)
				cobra.CheckErr(err)

				message := result.Choices[0].Message.Content
				fmt.Print(message[prevLen:])

				prevLen = len(message)
			}
		}

		fmt.Printf("\n\nTotal tokens: %d | Prompt: %d | Completion: %d\n", result.Usage.TotalTokens, result.Usage.PromptTokens, result.Usage.CompletionTokens)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().IntP(utils.MaxTokensKey, "m", 1000, "Maximum number of tokens to be used per request. Defaults to config value.")
	searchCmd.Flags().IntP(utils.TemperatureKey, "t", 0, "The amount of randomness in the response. Between 0 and 2.")
	searchCmd.Flags().IntP(utils.TopKKey, "K", 0, "Number of tokens to consider when generating tokens, lower values result in higher probability tokens being used. Between 0 and 2048.")
	searchCmd.Flags().IntP(utils.TopPKey, "P", 0, "Nucleus sampling. Probability cutoff for token selection, lower values result in higher probability tokens being used. Between 0 and 1.")
	searchCmd.Flags().IntP(utils.FrequencyPenaltyKey, "f", 0, "How much to penalize token reuse. 1 is no penalty. Between 0 and 1.")
	searchCmd.Flags().IntP(utils.PresencePenaltyKey, "p", 0, "How much to penalize existing tokens. Between -2 and 2.")
}