package utils

import (
	"github.com/spf13/viper"
)

func ResetConfig() {
	viper.Set(AdditionalInstructionsKey, "")
	viper.Set(ApiKeyKey, "")
	viper.Set(MaxTokensKey, "1000")
	viper.Set(ModelKey, "sonar-small-online")
	viper.WriteConfig()
}
