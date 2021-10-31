package http

const (
	HeaderAuth                        = "Authorization"
	ContextUserIDKey                  = "user_id"
	GeneralSuccessMessage             = "operation successful"
	AuthorizationNotExistErrorMessage = "Authorization header is required"
)

type GeneralAPIResponse struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func BuildAPIResponseError(statusCode int, err error) GeneralAPIResponse {
	return GeneralAPIResponse{
		Status: statusCode,
		Error:  err.Error(),
	}
}

func BuildAPIResponseSuccess(statusCode int, data interface{}) GeneralAPIResponse {
	return GeneralAPIResponse{
		Status: statusCode,
		Data:   data,
	}
}
