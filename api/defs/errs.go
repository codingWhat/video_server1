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
)