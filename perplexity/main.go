package perplexity

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Client struct {
	ApiKey   string
	Messages []Message
	Payload  map[string]any
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Result struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Created int    `json:"created"`
	Error   struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    int    `json:"code"`
	} `json:"error"`
	Usage struct {
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
	} `json:"choices"`
	SearchResults []struct {
		Date  string `json:"date"`
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"search_results"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		ApiKey:   apiKey,
		Messages: []Message{},
		Payload:  make(map[string]any),
	}
}

func (c *Client) SetPayload(key string, value any) {
	c.Payload[key] = value
}

func (c *Client) AppendMessage(role string, content string) {
	c.Messages = append(c.Messages, Message{role, content})
}

func (c *Client) getResponse() (*http.Response, error) {
	url := "https://api.perplexity.ai/chat/completions"

	payload := c.Payload
	payload["messages"] = c.Messages

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return response, err
	}
	if response.StatusCode != 200 {
		defer response.Body.Close()
		result := &Result{}

		err = json.NewDecoder(response.Body).Decode(result)
		if err != nil {
			return nil, err
		}
		if result.Error.Message == "" {
			result.Error.Message = "Unknown error"
		}

		return nil, errors.New(response.Status + ": " + result.Error.Message)
	}

	return response, err
}

func (c *Client) MakeRequest() (*Result, error) {
	response, err := c.getResponse()
	if err != nil {
		return nil, err
	}

	result := &Result{}

	err = json.NewDecoder(response.Body).Decode(result)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) MakeStreamedRequest(callback func(string)) (*Result, error) {
	response, err := c.getResponse()
	if err != nil {
		return nil, err
	}

	result := &Result{}
	scanner := bufio.NewScanner(response.Body)
	prevLen := 0

	for scanner.Scan() {
		bytes := scanner.Bytes()
		minLen := len("data: ")

		if len(bytes) < minLen {
			continue
		}

		data := bytes[minLen:]

		err := json.Unmarshal(data, result)
		if err != nil {
			return nil, err
		}

		message := result.Choices[0].Message.Content
		callback(message[prevLen:])

		prevLen = len(message)
	}

	defer response.Body.Close()
	return result, nil
}
