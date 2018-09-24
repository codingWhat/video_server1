package dbops

import (
	"strconv"
	"video_server/api/defs"
	"database/sql"
	"sync"
	"log"
)

// 写入session
func InsertSession(sid string,ttl int64, login_name string) error{
	ttlstr:=strconv.FormatInt(ttl,10)
	stmtInsert,err:=dbConn.Prepare("insert into sessions (session_id,TTL,login_name) values(? ,? , ?)")
	if err!=nil{
		return err
	}
	_,err=stmtInsert.Exec(sid,ttlstr,login_name)
	if err!=nil{
		return err
	}
	defer stmtInsert.Close()
	return nil
}

// 取回session
func RetrieveSession(sid string) (*defs.SimpleSession,error){
	ss:=&defs.SimpleSession{}
	stmtSelect,err:=dbConn.Prepare("select ttl,login_name from sessions where session_id = ?")

	var ttl string
	var uname string
	stmtSelect.QueryRow(sid).Scan(&ttl,&uname)
	if err!=nil && err!=sql.ErrNoRows{
		return nil,err
	}

	if res,err:=strconv.ParseInt(ttl,10,64);err!=nil{
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil,err
	}

	defer stmtSelect.Close()
	return ss,nil
}

// List Session
func RetrieveAllSessions() (*sync.Map,error){
	m:=&sync.Map{}
	stmtSelect,err:=dbConn.Prepare("select * from sessions")
	if err!=nil{
		log.Printf("%s",err)
		return nil,err
	}

	rows,err:=stmtSelect.Query()
	if err!=nil{
		log.Printf("%s",err)
		return nil,err
	}

	for rows.Next(){
		var id string
		var ttlstr string
		var login_name string
		if err:=rows.Scan(&id,&ttlstr,&login_name);err!=nil{
			log.Printf("retrive sessions error: %s",err)
			break
		}

		if ttl,err1:=strconv.ParseInt(ttlstr,10,64);err1!=nil{
			ss:=&defs.SimpleSession{
				Username:login_name,
				TTL:ttl,
			}
			m.Store(id,ss)
			log.Printf("session id: %s, ttl: %d",id,ss.TTL)
		}
	}
	return m,nil
}

// Delete Session
func DeleteSession(sid string) error {
	stmtSelect,err:=dbConn.Prepare("delete from sessions where session_id = ?")
	if err!=nil{
		log.Printf("%s",err)
		return err
	}

	if _,err:=stmtSelect.Query(sid);err!=nil{
		return err
	}
	return nil
}