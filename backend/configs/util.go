package configs

import "github.com/spf13/viper"

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
