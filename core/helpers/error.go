package helpers

type ErrorType int

const (
	BodyPropInvalid ErrorType = iota
	QueriesInvalid
	InternalServer
)

var ErrorMessages = map[ErrorType]string{
	BodyPropInvalid: "body properties invalid",
	QueriesInvalid:  "queries invalid",
	InternalServer:  "internal server error",
}

type ErrorResponse struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type GetErrorResponseArg struct {
	StatusCode int
	ErrorType  ErrorType
}

func GetErrorResponse(GetErrorResponseArg GetErrorResponseArg) ErrorResponse {
	description := ErrorMessages[GetErrorResponseArg.ErrorType]

	return ErrorResponse{
		Code:        GetErrorResponseArg.StatusCode,
		Description: description,
	}
}
