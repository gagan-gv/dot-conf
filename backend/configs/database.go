package configs

import "dot_conf/constants"

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getConfigFromEnvOrDefault(constants.DbHost, constants.HOST).(string),
		Port:     getConfigFromEnvOrDefault(constants.DbPort, constants.PORT).(int),
		User:     getConfigFromEnvOrDefault(constants.DbUser, constants.USER).(string),
		Password: getConfigFromEnvOrDefault(constants.DbPassword, constants.PASSWORD).(string),
		DbName:   getConfigFromEnvOrDefault(constants.DbName, constants.DB_NAME).(string),
		SslMode:  constants.SSLMODE,
	}
}
