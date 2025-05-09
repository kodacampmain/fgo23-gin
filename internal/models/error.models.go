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

var InternalServerErrorCode string = "INTERNAL_ERROR"
var DataNotFoundCode string = "DATA_NOT_FOUND"
var RegisterFailedCode string = "REGISTER_FAILED"
var InvalidUsernamePasswordCode string = "INVALID_USERNAME_PASSWORD"
