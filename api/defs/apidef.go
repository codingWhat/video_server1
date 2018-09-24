package defs

// requests
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id string // 因为是UUID
	AuthorId int
	Name string
	DisplayCtime string // 显示创建时间
}

// 评论
type Comment struct {
	Id string
	VideoId string
	Author string
	Content string
}

// Session会话
type SimpleSession struct {
	Username string // login name
	TTL int64
}