package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

// 添加一个用户信息
func AddUserCredential(loginName string, pwd string) error {
	// 千万不要用加号来连接query的各个部分，容易造成撞库和遭受攻击
	// Prepare是预编译sql
	stmtInsert, err := dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtInsert.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtInsert.Close() // 有性能损耗

	return nil
}

// 通过传入的登录名获取其密码
func GetUserCredential(loginName string) (string, error) {
	stmtSelect, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("query user pwd error:%s", err)
		return "", err
	}
	var pwd string
	// 把命中的扫描结果放到pwd里面
	err = stmtSelect.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtSelect.Close() // 有性能损耗

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDelete, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("Delete user error:%s", err)
		return err
	}

	_, err = stmtDelete.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	defer stmtDelete.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// creatime->db->
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // 输出M D y,HH:MM:SS
	stmtInsert, err := dbConn.Prepare("insert into video_info (id,author_id, name, display_ctime) values (?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtInsert.Exec(vid, aid, name, ctime)
	if err!=nil{
		return nil,err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtInsert.Close()

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtSelect, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id=?")

	var aid int
	var dct string
	var name string

	err = stmtSelect.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtSelect.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDelete, err := dbConn.Prepare("delete from video_info where id=?")
	if err != nil {
		return err
	}
	_, err = stmtDelete.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDelete.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtInsert, err := dbConn.Prepare("insert into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtInsert.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtInsert.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtSelect, err := dbConn.Prepare(`select comments.id,users.login_name,comments.content from comments
			inner join users on comments.author_id = users.id
			where comments.video_id = ? and comments.time > from_unixtime(?) and comments.time <=from_unixtime(?)`)
	var res []*defs.Comment

	rows, err := stmtSelect.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, nil
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtSelect.Close()
	return res, nil
}
