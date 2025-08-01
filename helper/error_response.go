package helper

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

func (e *ErrorResponse) ValidateRequestErrorCode() *ErrorResponse {
	e.Code = 0
	return e
}

func (e *ErrorResponse) ValidateServiceErrorCode() *ErrorResponse {
	e.Code = 1
	return e
}
