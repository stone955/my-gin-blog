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

	g.GET("/tags", GetTags)
	g.POST("/tags", AddTag)
	g.PUT("/tags/:id", EditTag)
	g.DELETE("/tags/:id", DeleteTag)

	return r
}
