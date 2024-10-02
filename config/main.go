package config

import (
	"github.com/japelsin/pplx/constants"
	"github.com/kirsle/configdir"
	"github.com/spf13/viper"
)

type config struct{}

func Init() error {
	configPath := configdir.LocalConfig()

	viper.AddConfigPath(configPath)
	viper.SetConfigName("pplx")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		if err == err.(viper.ConfigFileNotFoundError) {
			return Reset()
		}

		return err
	}

	return nil
}

func GetApiKey() string {
	return viper.GetString(constants.API_KEY_KEY)
}

func SetApiKey(value string) {
	viper.Set(constants.API_KEY_KEY, value)
}

func GetMaxTokens() int {
	return viper.GetInt(constants.MAX_TOKENS_KEY)
}

func SetMaxTokens(value int) {
	viper.Set(constants.MAX_TOKENS_KEY, value)
}

func GetModel() string {
	return viper.GetString(constants.MODEL_KEY)
}

func SetModel(value string) {
	viper.Set(constants.MODEL_KEY, value)
}

func GetSystemPrompt() string {
	return viper.GetString(constants.SYSTEM_PROMPT_KEY)
}

func SetSystemPrompt(value string) {
	viper.Set(constants.SYSTEM_PROMPT_KEY, value)
}

func GetStream() bool {
	return viper.GetBool(constants.STREAM_KEY)
}

func SetStream(value bool) {
	viper.Set(constants.STREAM_KEY, value)
}

func Save() error {
	return viper.WriteConfig()
}

func Reset() error {
	SetApiKey("")
	SetMaxTokens(constants.DEFAULT_MAX_TOKENS)
	SetModel(constants.DEFAULT_MODEL)
	SetSystemPrompt("")
	SetStream(true)

	return viper.SafeWriteConfig()
}
