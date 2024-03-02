package utils

import (
	"fmt"

	"github.com/kirsle/configdir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig() {
	path := configdir.LocalConfig()

	viper.AddConfigPath(path)
	viper.SetConfigName("pplx")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			ResetConfig()
			viper.SafeWriteConfig()
		} else {
			cobra.CheckErr(err)
		}
	}

	if viper.Get(ApiKeyKey) == "" {
		fmt.Println("API key not set")
		fmt.Println("Enter your Perplexity API key to get started")

		apiKey, err := Prompt("API key")
		cobra.CheckErr(err)

		viper.Set(ApiKeyKey, apiKey)
		viper.WriteConfig()
	}
}

func ResetConfig() {
	viper.Set(ApiKeyKey, "")
	viper.Set(MaxTokensKey, 1000)
	viper.Set(ModelKey, "sonar-small-online")
	viper.Set(TemperatureKey, 0.7)

	viper.WriteConfig()
}

func UpdateConfigValue(key string, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()

	cobra.CheckErr(err)
}
