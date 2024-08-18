package dto

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	DevMessage string      `json:"devMessage"`
	Timestamp  int64       `json:"timestamp"`
	Data       interface{} `json:"data,omitempty"`
}
