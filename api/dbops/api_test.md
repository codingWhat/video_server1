#### 通过postman工具测试api
```markdown
1. 创建用户
Request:
POST /user HTTP/1.1
Host: localhost:8000
Content-Type: application/json
{
	"user_name":"wangpengcheng",
	"pwd":"888"
}

Response：
{"success":true,"session_id":"0b465213-8d03-4afb-8aad-93d14aea70a7"}
```