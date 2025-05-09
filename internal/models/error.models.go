package models

type ErrorResponse struct {
	Error *ErrorResponseDetail `json:"error"`
}

type ErrorResponseDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Status  int    `json:"status"`
}
