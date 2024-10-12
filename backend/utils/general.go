package utils

import (
	"dot_conf/constants"
	"encoding/base64"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func ConvertJsonToString(j any, backup string) string {
	response, err := json.Marshal(j)
	if err != nil {
		log.Error("Error in marshalling response", err)
		if backup != constants.Empty {
			return backup
		}
		return constants.MarshallErrorResponse
	}

	return string(response)
}

func ConvertToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case float64: // if the number is stored as a float
		return int(v)
	case string:
		intValue, err := strconv.Atoi(v)
		if err != nil {
			log.Error("Error in converting to int", err)
			return 0
		}
		return intValue
	default:
		return 0
	}
}

func Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

func Decode(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Error("Error in decoding: ", err)
		return "", err
	}
	return string(decoded), nil
}
