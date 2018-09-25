package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// 把流控注入中间件层
type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}
// cc 表示流控大小,下面这种New函数类似于Java中的构造方法
func NewMiddleWareHandler(r *httprouter.Router,cc int) http.Handler{
	m:=middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if !m.l.GetConn(){
		sendErrorResponse(w,http.StatusTooManyRequests,"Too Many Requests") // 超过流控制返回状态吗：429
		return
	}

	m.r.ServeHTTP(w,r) //像api一样透传
	defer m.l.ReleaseConn()
}

func RegisterHandlers() *httprouter.Router{
	router:=httprouter.New()

	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)

	router.GET("/testpage",testPageHandler)
	return router
}

func main() {
	r:=RegisterHandlers()
	mh:=NewMiddleWareHandler(r,20)
	http.ListenAndServe(":9000",mh)

	// go install 然后启动bin/streamserver
	// 打开2个tab页访问下
	// http://localhost:9000/videos/shaonianxing.mp4
	// 用lsof -i:9000查看连接数

	// 访问上传文件,上传成功之后然后再访问这个上传的文件看看是否可以正常播放
	// http://localhost:9000/testpage
}
