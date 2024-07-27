package dto

import (
	"dot_conf/services"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	DevMessage string      `json:"devMessage"`
	Timestamp  int64       `json:"timestamp"`
	Data       interface{} `json:"data,omitempty"`
}

type Handler struct {
	CompanyService services.ICompanyService
}
