#### API设计：用户
```markdown
- 创建(注册)用户：URL:/user Method: POST,StatusCode:201,400,500
- 用户登录：URL:/user/:username Method: POST,StatusCode:200,400,500
- 获取用户基本信息：URL:/user/:username Method: GET,StatusCode:200,400,401,403,500
- 用户注销：URL:/user/:username Method: DELETE, StatusCode:204,400,401,403,500
```
---
#### API设计：用户资源
```markdown
- List all videos: URL:/user/:username/videos Method:GET,StatusCode:200,400,500,需要考虑分页的问题
- Get one video: URL:/user/username/videos/:vid-id Method:GET,StatusCode:200,400,500
- Delete one video: URL:/user/:username/videos/:vid-id Method:DELETE,StatusCode:204,400,401,403,500
```
---
#### API设计：评论
```markdown
- Show comments: URL:/videos/:vid-id/comments Method:GET,StatusCode:200,400,500
- Post a comment: URL:/videos/:vid-id/comments Method:POST,StatusCode:201,400,500
- Delete a comment: URL:/videos/:vid-id/comment/:comment-id Method:DELETE,StatusCode:204,400,401,403,500
```