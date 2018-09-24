package main

import (
	"net/http"
	"video_server/api/defs"
	"github.com/gin-gonic/gin/json"
	"io"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	// 序列化
	resStr,_:=json.Marshal(&errResp.Error)
	io.WriteString(w,string(resStr))
}

func sendNormalResponse(w http.ResponseWriter,resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}
