package dbops

import (
	"testing"
	"fmt"
)

// init(dblogin,truncate tables) -> run tests -> clear data(truncate tables)
// 顺序性的测试一个完整的增删改查操作

// 用TestMain函数来初始化所有test case
func TestMain(m *testing.M){
	clearTables()

	m.Run()

	clearTables()
}

// 清理表数据
func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

// 为了保证测试case的顺序必须用sub test
func TestUserWorkFlow(t *testing.T){
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Delete",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T){
	err:=AddUserCredential("wpc","123")
	if err!=nil{
		t.Errorf("Error of AddUser: %v",err)
	}
}

func testGetUser(t *testing.T){
	pwd,err:=GetUserCredential("wpc")
	if pwd!="123"||err!=nil{
		t.Errorf("Error of GetUser: %v",err)
	}
	fmt.Println("pwd:",pwd)
}

func testDeleteUser(t *testing.T){
	err:=DeleteUser("wpc","123")
	if err!=nil{
		t.Errorf("Error of DeleteUser: %v",err)
	}
}

// 测试被delete之后是否真的被删除了
func testRegetUser(t *testing.T){
	pwd,err:=GetUserCredential("wpc")
	if err!=nil{
		t.Errorf("Error of RegetUser: %v",err)
	}
	if pwd!=""{
		t.Errorf("Deleting user test failed")
	}
}