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

## 部署到 Docker
### 安装 Docker
````
# 卸载旧版本
yum remove -y docker \
docker-client \
docker-client-latest \
docker-common \
docker-latest \
docker-latest-logrotate \
docker-logrotate \
docker-selinux \
docker-engine-selinux \
docker-engine

# 设置 yum repository
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo

# 设置镜像仓库（可选）
vi /etc/docker/daemon.json

{
    "registry-mirrors": ["https://registry.docker-cn.com"]
}


# 安装并启动 docker
yum install -y docker-ce{-18.09.9} 
systemctl enable docker
systemctl start docker
systemctl stop docker
systemctl restart docker
````

### 编写 Dockerfile

### 构建镜像
+ 进入 my-gin-blog 根目录
````
docker build -t my-gin-blog-docker .
````

### 验证镜像
````
[root@localhost my-gin-blog]# docker images
REPOSITORY           TAG                 IMAGE ID            CREATED                  SIZE
golang               latest              272e3f68338f        Less than a second ago   803MB
my-gin-blog-docker   latest              80a0a9a255ac        16 seconds ago           1.11GB
````

### 创建并运行一个容器
````
[root@localhost my-gin-blog]# docker run -p 8080:8080 my-gin-blog-docker
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/stone955/my-gin-blog/router/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/stone955/my-gin-blog/router/api/v1.GetTags (4 handlers)
[GIN-debug] GET    /api/v1/tags/:id          --> github.com/stone955/my-gin-blog/router/api/v1.GetTag (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/stone955/my-gin-blog/router/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/stone955/my-gin-blog/router/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/stone955/my-gin-blog/router/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/stone955/my-gin-blog/router/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/stone955/my-gin-blog/router/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/stone955/my-gin-blog/router/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/stone955/my-gin-blog/router/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/stone955/my-gin-blog/router/api/v1.DeleteArticle (4 handlers)
2020/01/05 09:48:30 Actual pid is 1
````

### 验证容器实例
````
[root@localhost ~]# docker ps
CONTAINER ID        IMAGE                COMMAND             CREATED             STATUS              PORTS                    NAMES
5b25ed2ed214        my-gin-blog-docker   "./my-gin-blog"     15 minutes ago      Up 15 minutes       0.0.0.0:8080->8080/tcp   romantic_sutherland
````

### 删除镜像
````
# 查看 container
docker ps
# 查看 关联的容器
docker ps -a
# 停止运行中的 container
docker stop 5b25ed2ed214
# 删除 container
docker rm 5b25ed2ed214
# 删除 image
docker rmi 80a0a9a255ac
````

### mysql 容器化

#### 拉取 mysql 镜像
````
# 如果失败就多试几次或配置国内镜像仓库
[root@localhost my-gin-blog]# docker pull mysql
Using default tag: latest
latest: Pulling from library/mysql
Digest: sha256:e1b0fd480a11e5c37425a2591b6fbd32af886bfc6d6f404bd362be5e50a2e632
Status: Image is up to date for mysql:latest
docker.io/library/mysql:latest
````

#### 运行 mysql 容器
````
[root@localhost my-gin-blog]# docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql
46115d4847e53d030344e94c26d10c2167e348c6ed95fbdbb2f7a90b6b20c6aa
````