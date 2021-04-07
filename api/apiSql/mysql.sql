CREATE DATABASE api;
drop database if exists api;
create database api default character set utf8mb4 collate utf8mb4_unicode_ci;

use api;

DROP TABLE IF EXISTS user;
CREATE TABLE user (
    id int(10) unsigned not null AUTO_INCREMENT,
    username varchar(50) DEFAULT '' COMMENT '账号',
    password varchar(50) DEFAULT '' COMMENT '密码',
    head varchar(50) DEFAULT '' COMMENT '头像',
    address varchar(50) DEFAULT '' COMMENT '地址',
    primary key (id)
) ENGINE=InnoDB AUTO_INCREMENT=2 CHARSET = utf8;


# 设置默认值
