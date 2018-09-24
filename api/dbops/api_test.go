package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

// init(dblogin,truncate tables) -> run tests -> clear data(truncate tables)
// 顺序性的测试一个完整的增删改查操作

// 用TestMain函数来初始化所有test case
func TestMain(m *testing.M) {
	clearTables()

	m.Run()

	clearTables()
}

// 清理表数据
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

// 为了保证测试case的顺序必须用sub test
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("wpc", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("wpc")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
	fmt.Println("pwd:", pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("wpc", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

// 测试被delete之后是否真的被删除了
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("wpc")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

// 为了保证测试case的顺序必须用sub test
// 测试video对象的增，删，查的api
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddNewVideo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}
func testAddNewVideo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddNewVideo: %v", err)
	}
	tempvid = vi.Id
}
func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}
func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}
func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

// 为了保证测试case的顺序必须用sub test
// 测试Comment对象的增，查api
func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddNewComments)
	t.Run("ListComments", testListComments)

}
func testAddNewComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}
func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
