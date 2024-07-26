package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func generateAppKey() (string, error) {
	length := 16
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	hexString := hex.EncodeToString(bytes)

	apiKey := fmt.Sprintf("%s-%s-%s-%s", hexString[0:4], hexString[4:8], hexString[8:12], hexString[12:16])

	return apiKey, nil
}
