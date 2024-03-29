package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"io/ioutil"
	"video_server/api/defs"
	"encoding/json"
	"video_server/api/dbops"
	"video_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w, "Create User Handler") // 演示使用
	res,_:=ioutil.ReadAll(r.Body)
	ubody:=&defs.UserCredential{}
	// json.Unmarshal,反序列化的过程，把json转为struct
	if err:=json.Unmarshal(res,ubody) ;err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if err:=dbops.AddUserCredential(ubody.UserName,ubody.Pwd); err!=nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	//
	// 创建一个session
	id:=session.GenerateNewSessionId(ubody.UserName)
	su:=&defs.SignedUp{Success:true,SessionId:id}

	if resp,err:=json.Marshal(su);err!=nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else {
		sendNormalResponse(w,string(resp),201)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
