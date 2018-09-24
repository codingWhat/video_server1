package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler{
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	// check session
	validateUserSession(r)

	m.r.ServeHTTP(w,r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	mh:=NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

// golang中每一个goroutine只占4k内存
// listen->RegisterHandlers->handlers
// handlers->需要校验：1.校验request请求是否合法，2.校验user是否已注册的合法用户->然后business logic业务逻辑处理->response
/**
校验部分：
	1. data model
	2. error handling
*/

// 项目执行流程
//main->middleware(鉴权等)->defs(message,err)->handlers->dbops->response
