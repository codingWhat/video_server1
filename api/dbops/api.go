package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 添加一个用户信息
func AddUserCredential(loginName string,pwd string) error{
	// 千万不要用加号来连接query的各个部分，容易造成撞库和遭受攻击
	// Prepare是预编译sql
	stmtInsert,err:=dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")
	defer stmtInsert.Close() // 有性能损耗
	if err!=nil{
		log.Printf("%s",err)
		return err
	}
	stmtInsert.Exec(loginName,pwd)
	//stmtInsert.Close() // 当然也可以加上defer,但是defer会有性能上的损耗的，如果是defer stmtInsert.Close()就直接卸载stmtInsert下一行即可

	return nil
}

// 通过传入的登录名获取其密码
func GetUserCredential(loginName string) (string,error){
	stmtSelect,err:=dbConn.Prepare("select pwd from users where login_name = ?")
	defer stmtSelect.Close() // 有性能损耗
	if err!=nil{
		log.Printf("query user pwd error:%s",err)
		return "",err
	}
	var pwd string
	// 把命中的扫描结果放到pwd里面
	stmtSelect.QueryRow(loginName).Scan(&pwd)

	return  pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	stmtDelete,err:=dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	defer stmtDelete.Close()
	if err!=nil{
		log.Printf("Delete user error:%s",err)
		return err
	}
	stmtDelete.Exec(loginName,pwd)

	return nil
}
