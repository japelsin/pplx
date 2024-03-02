package utils

const (
	ApiKeyKey           = "api_key"
	FrequencyPenaltyKey = "frequency_penalty"
	MaxTokensKey        = "max_tokens"
	ModelKey            = "model"
	PresencePenaltyKey  = "presence_penalty"
	QueryKey            = "query"
	TemperatureKey      = "temperature"
	TopKKey             = "top_k"
	TopPKey             = "top_p"
)

var AvailableModels = []string{"sonar-small-chat", "sonar-small-online", "sonar-medium-chat", "sonar-medium-online", "llama-2-70b-chat", "codellama-34b-instruct", "codellama-70b-instruct", "mistral-7b-instruct", "mixtral-8x7b-instruct"}
