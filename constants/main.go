package constants

const (
	API_KEY_KEY               = "api_key"
	DEFAULT_MAX_TOKENS        = 1000
	DEFAULT_MODEL             = "sonar"
	FREQUENCY_PENALTY_KEY     = "frequency_penalty"
	MAX_TOKENS_KEY            = "max_tokens"
	MODEL_KEY                 = "model"
	PRESENCE_PENALTY_KEY      = "presence_penalty"
	SEARCH_DOMAIN_FILTER_KEY  = "search_domain_filter"
	SEARCH_RECENCY_FILTER_KEY = "search_recency_filter"
	STREAM_KEY                = "stream"
	SYSTEM_PROMPT_KEY         = "system_prompt"
	TEMPERATURE_KEY           = "temperature"
	TOP_K_KEY                 = "top_k"
	TOP_P_KEY                 = "top_p"

	// Closed beta
	// RETURN_CITATIONS_KEY         = "return_citations"
	// RETURN_IMAGES_KEY            = "return_images"
	// RETURN_RELATED_QUESTIONS_KEY = "return_related_questions"
)

var (
	AVAILABLE_MODELS       = []string{DEFAULT_MODEL, "sonar-pro", "sonar-pro", "sonar-pro", "sonar-reasoning-pro"}
	CONFIG_KEYS            = []string{API_KEY_KEY, MAX_TOKENS_KEY, MODEL_KEY, STREAM_KEY, SYSTEM_PROMPT_KEY}
	SEARCH_RECENCY_FILTERS = []string{"month", "week", "day", "hour"}
)
