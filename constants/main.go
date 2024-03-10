package constants

const (
	DEFAULT_MODEL         = "sonar-small-online"
	DEFAULT_MAX_TOKENS    = 1000
	API_KEY_KEY           = "api_key"
	MAX_TOKENS_KEY        = "max_tokens"
	MODEL_KEY             = "model"
	STREAM_KEY						= "stream"
	SYSTEM_PROMPT_KEY     = "system_prompt"
	FREQUENCY_PENALTY_KEY = "frequency_penalty"
	PRESENCE_PENALTY_KEY  = "presence_penalty"
	TEMPERATURE_KEY       = "temperature"
	TOP_K_KEY             = "top_k"
	TOP_P_KEY             = "top_p"
)

var (
	AVAILABLE_MODELS = []string{"sonar-small-chat", DEFAULT_MODEL, "sonar-medium-chat", "sonar-medium-online", "llama-2-70b-chat", "codellama-34b-instruct", "codellama-70b-instruct", "mistral-7b-instruct", "mixtral-8x7b-instruct"}
	CONFIG_KEYS      = []string{API_KEY_KEY, MAX_TOKENS_KEY, MODEL_KEY, STREAM_KEY, SYSTEM_PROMPT_KEY}
)
