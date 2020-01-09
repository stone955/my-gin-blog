package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/setting"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	g := r.Group("/api")

	// 标签
	g.GET("/tags", GetTags)
	g.GET("/tags/:id", GetTag)
	g.POST("/tags", AddTag)
	g.PUT("/tags/:id", EditTag)
	g.DELETE("/tags/:id", DeleteTag)

	// 文章
	g.GET("/articles", GetArticles)
	g.GET("/articles/:id", GetArticle)
	g.POST("/articles", AddArticle)
	g.PUT("/articles/:id", EditArticle)
	g.DELETE("/articles/:id", DeleteArticle)

	return r
}
