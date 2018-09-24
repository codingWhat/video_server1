package defs

import "net/http"

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC:http.StatusNotFound,Error:Err{Error:"Request body is not corrent",ErrorCode:"001"}}
	ErrorNotAuthUser = ErrorResponse{HttpSC:http.StatusUnauthorized,Error:Err{Error:"User authentication failed",ErrorCode:"002"}}
	// 数据库操作异常
	ErrorDBError = ErrorResponse{HttpSC:500,Error:Err{Error:"DB ops failed",ErrorCode:"003"}}
	// 内部服务异常
	ErrorInternalFaults = ErrorResponse{HttpSC:500,Error:Err{Error:"Internal service error",ErrorCode:"004"}}
)