package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", homeHandler)
	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)

	// api 透传
	router.POST("/api", apiHandler)
	// api proxy
	router.POST("upload/:vid-id", proxyHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}

// cross origin resource sharing:CORS 跨域访问
/**
# 浏览器访问地址是：
http://127.0.0.1:8080/upload/:vid-id
# 实际后台地址是：
http://127.0.0.1:9000/upload/:vid-id
# 上面2个url不再同一个域中，浏览器会因为安全隐患问题阻止访问，俗称跨域访问，那么我们就需要用到proxy功能
*/
