package main

import (
	"github.com/gin-gonic/gin"
	api "github.com/stone955/my-gin-blog/api"
	"github.com/stone955/my-gin-blog/pkg/setting"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	g := r.Group("/api")

	g.GET("/tags", api.GetTags)
	g.POST("/tags", api.AddTag)
	g.PUT("/tags/:id", api.EditTag)
	g.DELETE("/tags/:id", api.DeleteTag)

	return r
}
