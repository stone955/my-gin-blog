package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/export"
	"github.com/stone955/my-gin-blog/pkg/upload"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	_ "github.com/stone955/my-gin-blog/docs"
	"github.com/stone955/my-gin-blog/middleware/jwt"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/router/api"
	"github.com/stone955/my-gin-blog/router/api/v1"
)

func Register() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerCfg.RunMode)

	// Jwt
	r.GET("/auth", api.GetAuth)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// upload
	r.POST("/upload", api.UploadImage)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export/excel", http.Dir(export.GetExcelFullPath()))

	g := r.Group("/api/v1")

	// 添加中间件
	g.Use(jwt.JWT())

	// 标签
	g.GET("/tags", v1.GetTags)
	g.GET("/tags/:id", v1.GetTag)
	g.POST("/tags", v1.AddTag)
	g.PUT("/tags/:id", v1.EditTag)
	g.DELETE("/tags/:id", v1.DeleteTag)
	g.POST("/tags/export", v1.ExportTag)
	g.POST("/tags/import", v1.ImportTag)

	// 文章
	g.GET("/articles", v1.GetArticles)
	g.GET("/articles/:id", v1.GetArticle)
	g.POST("/articles", v1.AddArticle)
	g.PUT("/articles/:id", v1.EditArticle)
	g.DELETE("/articles/:id", v1.DeleteArticle)

	return r
}
