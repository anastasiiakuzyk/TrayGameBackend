package utils

type ConnectRequest struct {
	UUID string `json:"uuid"`
}

type ErrorPackage struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

func NewErrorPackage(err string) *ErrorPackage {
	return &ErrorPackage{
		Type:  "error",
		Error: err,
	}
}

type Response struct {
}

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code uint, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

type StartResponse struct {
	Code uint   `json:"code"`
	UUID string `json:"uuid"`
}

func NewStartResponse(code uint, uuid string) *StartResponse {
	return &StartResponse{
		Code: code,
		UUID: uuid,
	}
}
