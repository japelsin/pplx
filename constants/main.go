package constants

const (
	API_KEY_KEY              = "api_key"
	DEFAULT_MAX_TOKENS       = 1000
	DEFAULT_MODEL            = "sonar"
	DEFAULT_SYSTEM_PROMPT    = "Be precise and concise."
	MAX_TOKENS_CONFIG_KEY    = "max_tokens"
	MODEL_CONFIG_KEY         = "model"
	STREAM_CONFIG_KEY        = "stream"
	SYSTEM_PROMPT_CONFIG_KEY = "system_prompt"
)

var (
	AVAILABLE_MODELS       = []string{DEFAULT_MODEL, "sonar-pro", "sonar-pro", "sonar-pro", "sonar-reasoning-pro"}
	CONFIG_KEYS            = []string{API_KEY_KEY, MAX_TOKENS_CONFIG_KEY, MODEL_CONFIG_KEY, STREAM_CONFIG_KEY, SYSTEM_PROMPT_CONFIG_KEY}
	SEARCH_RECENCY_FILTERS = []string{"month", "week", "day", "hour"}
)
