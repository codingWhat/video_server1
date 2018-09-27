package main

import (
	"bytes"
	"github.com/gin-gonic/gin/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 代理：转发业务请求，渲染
// proxy
/*
	比如前端访问http://127.0.0.1:8000/upload,但是后台真正去访问的是
	http://127.0.0.1:9000/upload,这叫proxy
*/

// api
/**
	比如前端访问http://127.0.0.1:8000/api/v1/pods,后台会取出来封装成
{
	"url": "",
	"method": "",
	"message": ""
}
组装到去访问后端的apiserver
*/

// API 透传模块实现

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	// api一般分为GET POST DELETE
	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("%s", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("%s", err)
			return
		}
		normalResponse(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", b.Url, nil)
		req.Header = r.Header
		resp, err := httpClient.Do(req)
		if err != nil {
			log.Printf("%s", err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Bad api request")
	}
}

// 真正的透传，把*http.Response写到responsewriter里面
func normalResponse(w http.ResponseWriter, r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w, string(res))
}
