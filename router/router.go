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

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	return r
}
