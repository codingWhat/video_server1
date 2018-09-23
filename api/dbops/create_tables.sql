-- 创建video_server数据库
create database video_server;

-- 创建users用户表
create table users(
  id int(10) unsigned primary key not null AUTO_INCREMENT,
  login_name varchar(64) unique key,
  pwd text not null
);
-- 创建video_info表
create table video_server(
  id varchar(64) primary key not null ,
  author_id int(10) unsigned,
  name text,
  display_ctime text,
  create_time datetime default current_timestamp
);
-- 创建comments表
create table comments(
  id varchar(64) primary key not null ,
  video_id varchar(64),
  author_id int(10) unsigned,
  content text,
  time datetime default current_timestamp
);
-- 创建sessions表
create table sessions(
  --session_id tinytext primary key not null ,有问题！
  session_id varchar(64) primary key not null ,
  TTL tinytext,
  login_name varchar(64)
);

-- 命令行操作数据库
-- 切换数据库
use video_server;
-- 显示当前数据库下的所有表
show tables;
-- 查看users表
describe users;
-- 查看video_info表
describe video_info
-- 查看comments表
describe comments
-- 查看sessions表
describe sessions
