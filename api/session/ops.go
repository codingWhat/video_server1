package session

import (
	"sync"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
	"time"
)

// session是否过期，session需要存储
// 可以用redis作为web的缓存，但是没必要，要综合考量
// 也可以使用golang 内置提供的缓存机制，sync.Map是缓存，1.9才支持

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}
// 没有返回值，因为是在内存操作的，不需要返回值
func LoadSessionsFromDB(){
	r,err:=dbops.RetrieveAllSessions()
	if err!=nil{
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss:=value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

// 既要往本身cache写也要往db中写
func GenerateNewSessionId(unname string) string {
	id,_:=utils.NewUUID()
	//ct:=time.Now().Unix()/1000000 // 毫秒表示
	ct:=nowInMilli()
	ttl:=ct + 30*60*1000// Severside session valid time: 30min

	ss:=&defs.SimpleSession{
		Username:unname,
		TTL:ttl,
	}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,unname)
	return id
}
// 过期则返回("",true)
// 没有过期则返回(sessionId,false)
func IsSessionExpired(sid string) (string,bool){
	ss,ok:=sessionMap.Load(sid)
	if ok{
		//ct:=time.Now().UnixNano()/1000000
		ct:=nowInMilli()
		if ss.(*defs.SimpleSession).TTL <ct {
			// delete expired session
			deleteExpiredSession(sid)
			return "",true // 过期了返回true和空session
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}
