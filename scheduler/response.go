package main

import (
	"io"
	"net/http"
)

//sc:statuccode, resp "msg"
func sendResponse(w http.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
