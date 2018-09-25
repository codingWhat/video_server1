package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"html/template"
)

// 测试
func testPageHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params){
	t,_:=template.ParseFiles("./videos/upload.html")
	t.Execute(w,nil)
}

// 将本地的mp4文件以stream流的形式传递到client端(如浏览器端)
func streamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	vid:=p.ByName("vid-id")
	vl:=VIDEO_DIR+vid

	// 下一步打开这个video
	video,err:=os.Open(vl)
	defer video.Close()
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)
}

// 将client端的mp4文件上传到server端
/**
第一步：校验传入的文件的大小
。。。
 */
func uploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	// 限制io.Reader最大能读的文件大小
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err:=r.ParseMultipartForm(MAX_UPLOAD_SIZE);err!=nil{
		fmt.Printf("Error when try to open file: %v",err)
		sendErrorResponse(w,http.StatusBadRequest,"File is to big,Bad Request")
		return
	}

	file,_,err:=r.FormFile("file") // <form name="file"
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	data,err:=ioutil.ReadAll(file)
	if err!=nil{
		log.Printf("Read file error: %v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	fn:=p.ByName("vid-id")
	// 上传文件就是写入文件
	err = ioutil.WriteFile(VIDEO_DIR+fn,data,0666)
	if err!=nil{
		log.Printf("Write file error: %v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"uploaded successfully")
}
