package handler

type ErrorCode int

const (
	SuccessCode ErrorCode = 0
	SuccessMsg  string    = ""

	JsonDecodeErrorCode ErrorCode = 1
	JsonDecodeErrorMsg  string    = "JSON decode error:"

	EnvErrorCode ErrorCode = 2
	EnvErrorMsg  string    = "TI Environment error:"

	ParamErrorCode ErrorCode = 3
	ParamErrorMsg  string    = "Invalid parameter:"
)
