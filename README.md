## 新建工程
## 创建数据库
+ 标签表
```
CREATE TABLE `t_blog_tag` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
  `created_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '创建时间',
  `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
  `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
  `modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` INT(10) UNSIGNED DEFAULT '0',
  `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```
+ 文章表
````
CREATE TABLE `t_blog_article` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_id` INT(10) UNSIGNED DEFAULT '0' COMMENT '标签ID',
  `title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
  `desc` VARCHAR(255) DEFAULT '' COMMENT '简述',
  `content` TEXT,
  `created_on` INT(11) DEFAULT NULL,
  `created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
  `modified_on` INT(10) UNSIGNED DEFAULT '0' COMMENT '修改时间',
  `modified_by` VARCHAR(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` INT(10) UNSIGNED DEFAULT '0',
  `state` TINYINT(3) UNSIGNED DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='文章管理';
````
+ 认证表
````
CREATE TABLE `t_blog_auth` (
  `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(50) DEFAULT '' COMMENT '账号',
  `password` VARCHAR(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;

INSERT INTO `t_blog_auth` (`username`, `password`) VALUES ('admin', '123456');
````

## 编译
go build -o my-gin-blog main.go

## 集成Swagger
### 安装swag
````
go get github.com/swaggo/swag/cmd/swag
````
````
-- 验证安装是否成功
swag -v
swag version v1.6.4
````
### 安装gin-swagger
````
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/gin-swagger/swaggerFiles
````

### 添加注释
（略）
### 生成
````
-- 进入项目根目录
[root@localhost my-gin-blog]# swag init
````