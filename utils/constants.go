package utils

const (
	AdditionalInstructionsKey = "additional_instructions"
	ApiKeyKey                 = "api_key"
	MaxTokensKey              = "max_tokens"
	ModelKey                  = "model"
	TemperatureKey            = "temperature"
	TopPKey                   = "top_p"
	TopKKey                   = "top_k"
	PresencePenaltyKey        = "presence_penalty"
	FrequencyPenaltyKey       = "frequency_penalty"
)

var AvailableModels = []string{"sonar-small-chat", "sonar-small-online", "sonar-medium-chat", "sonar-medium-online", "llama-2-70b-chat", "codellama-34b-instruct", "codellama-70b-instruct", "mistral-7b-instruct", "mixtral-8x7b-instruct"}
