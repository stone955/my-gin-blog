package v1

import "github.com/gin-gonic/gin"

// GetArticles 查询所有文章
// curl localhost:8080/api/v1/articles
func GetArticles(c *gin.Context) {

}

// GetArticle 获取一个文章
// curl localhost:8080/api/v1/articles/:id
func GetArticle(c *gin.Context) {

}

// AddArticle 新建文章
// curl -X POST localhost:8080/api/v1/articles -d "{\"name\":\"golang\",\"state\":1,\"created_by\":\"admin\"}"
func AddArticle(c *gin.Context) {

}

// EditArticle 更新指定文章
// curl -X PUT localhost:8080/api/v1/articles/:id -d "{\"name\":\"golang\",\"state\":1,\"created_by\":\"admin\"}"
func EditArticle(c *gin.Context) {

}

// DeleteArticle 删除指定文章
// curl -X DELETE localhost:8080/api/v1/articles/:id
func DeleteArticle(c *gin.Context) {

}
