package main

import (
	"net/http"
	"video_server/api/session"
	"video_server/api/defs"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

// session校验
func validateUserSession(r *http.Request) bool{
	sid:=r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) ==0 {
		return false
	}
	uname,ok:=session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME,uname)
	return true
}

// 用户校验
func ValidateUser(w http.ResponseWriter,r *http.Request) bool{
	uname:=r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}
