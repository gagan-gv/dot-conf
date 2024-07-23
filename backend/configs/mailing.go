package configs

import "dot_conf/constants"

type MailingConfig struct {
	Host     string
	Port     int
	Password string
	From     string
}

func NewMailingConfig() *MailingConfig {
	return &MailingConfig{
		Host:     getConfigFromEnvOrDefault(constants.MailHost, "").(string),
		Port:     getConfigFromEnvOrDefault(constants.MailPort, 0).(int),
		Password: getConfigFromEnvOrDefault(constants.MailPassword, "").(string),
		From:     getConfigFromEnvOrDefault(constants.MailUser, "").(string),
	}
}
