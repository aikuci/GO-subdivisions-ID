package error

import "net/http"

var (
	StatusBadRequest = "400"
	StatusNotFound   = "404"

	StatusInternalServerError = "500"
)

type CustomErrorResponse struct {
	Message  string `json:"message,omitempty"`
	ErrCode  string `json:"code,omitempty"`
	HTTPCode int    `json:"http_code"`
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomErrorResponse) Error() string {
	return c.Message
}

func CustomError(message string, errCode string, httpCode int) *CustomErrorResponse {
	return &CustomErrorResponse{
		Message:  message,
		ErrCode:  errCode,
		HTTPCode: httpCode,
	}
}

func BadRequest(message string) *CustomErrorResponse {
	return &CustomErrorResponse{
		Message:  message,
		ErrCode:  StatusBadRequest,
		HTTPCode: http.StatusBadRequest,
	}
}

func RecordNotFound(message string) *CustomErrorResponse {
	return &CustomErrorResponse{
		Message:  message,
		ErrCode:  StatusNotFound,
		HTTPCode: http.StatusNotFound,
	}
}

func InternalServerError(message string) *CustomErrorResponse {
	return &CustomErrorResponse{
		Message:  message,
		ErrCode:  StatusInternalServerError,
		HTTPCode: http.StatusInternalServerError,
	}
}
