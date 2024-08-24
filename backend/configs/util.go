package configs

import (
	"dot_conf/constants"
	"github.com/spf13/viper"
)

const (
	defaultKey = "default-key"
)

func getConfigFromEnvOrDefault(configName string, defaultValue any) any {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return defaultValue
	}

	value := viper.Get(configName)

	if value == nil {
		return defaultValue
	}
	return value
}

func GetJwtSecretKey() []byte {
	return getConfigFromEnvOrDefault(constants.JwtSecretKey, defaultKey).([]byte)
}
