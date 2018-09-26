- 什么是scheduler
- 为什么需要scheduler
- scheduler通常做什么

#### scheduler包含什么
```markdown
- Restful的http server
- Timer
- 生产者/消费者模型下的task runner

```

#### 测试这个scheduler模块
```markdown
cd /media/wangfeng/deepin-f/GitKraken/github/src/video_server/scheduler
# 下面这个会在bin/目录下生成scheduler二进制文件，执行这个二进制文件即可启动这个scheduler模块
go install

wangfeng@wangfeng-PC:/media/wangfeng/deepin-f/GitKraken/github/bin/videos$ touch 123 1234 12345

# 启动
bin/scheduler

# 查询数据库
select * from video_del_rec;
# 清除表中的数据

# 我们通过postman来执行
GET http://127.0.0.1:9001/video-delete-record/123

```