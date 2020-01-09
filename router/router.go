package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/router/v1"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	g := r.Group("/api/v1")

	// 标签
	g.GET("/tags", v1.GetTags)
	g.GET("/tags/:id", v1.GetTag)
	g.POST("/tags", v1.AddTag)
	g.PUT("/tags/:id", v1.EditTag)
	g.DELETE("/tags/:id", v1.DeleteTag)

	// 文章
	g.GET("/articles", v1.GetArticles)
	g.GET("/articles/:id", v1.GetArticle)
	g.POST("/articles", v1.AddArticle)
	g.PUT("/articles/:id", v1.EditArticle)
	g.DELETE("/articles/:id", v1.DeleteArticle)

	return r
}
