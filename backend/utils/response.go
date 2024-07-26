package utils

import (
	"dot_conf/dto"
	"time"
)

func NewSuccessResponse(statusCode int, message, devMessage string, data interface{}) dto.Response {
	return dto.Response{
		StatusCode: statusCode,
		Message:    message,
		DevMessage: devMessage,
		Data:       data,
		Timestamp:  time.Now().UnixMilli(),
	}
}

func NewErrorResponse(statusCode int, message, devMessage string) dto.Response {
	return dto.Response{
		StatusCode: statusCode,
		Message:    message,
		DevMessage: devMessage,
		Timestamp:  time.Now().UnixMilli(),
	}
}
