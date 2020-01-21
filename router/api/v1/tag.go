package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/models"
	"github.com/stone955/my-gin-blog/pkg/app"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/pkg/util"
	"github.com/unknwon/com"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"time"
)

type Tag struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	CreatedBy string     `json:"created_by,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedBy string     `json:"updated_by,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedBy string     `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	State     int        `json:"state"`
}

// @Summary 新增文章标签
// @Produce json
// @Param name query string false "Name"
// @Param state query string false "State"
// @Success 200 {string} string "{"code": 200,"data": {},"msg": "ok"}"
// @Router /api/v1/tags [get]
// GetTags 查询所有标签
// curl 192.168.1.108:8080/api/v1/tags?token=
func GetTags(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		data = make([]Tag, 0)
	)

	query := make(map[string]interface{})

	name := c.Query("name")

	if name != "" {
		query["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		query["state"] = state
	}

	tags, err := models.GetTags(util.GetPage(c), setting.AppCfg.PageSize, query)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.Error, data)
		return
	}

	for _, tag := range tags {
		data = append(data, Tag{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt,
			CreatedBy: tag.CreatedBy,
			UpdatedAt: tag.UpdatedAt,
			UpdatedBy: tag.UpdatedBy,
			DeletedAt: tag.DeletedAt,
			DeletedBy: tag.DeletedBy,
			State:     tag.State,
		})
	}
	appG.Response(http.StatusOK, e.Ok, data)
}

// GetTag 查询指定标签
// curl 192.168.1.108:8080/api/tags/:id?token=
func GetTag(c *gin.Context) {

}

// AddTag 添加新标签
// curl -X POST 192.168.1.108:8080/api/v1/tags?token= -d "{\"name\":\"golang\",\"state\":1,\"created_by\":\"admin\"}"
func AddTag(c *gin.Context) {

	var (
		appG = app.Gin{C: c}
		data = Tag{}
	)

	if err := c.BindJSON(&data); err != nil {
		log.Printf("Error to bind json: %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}

	if err := validator.Valid(data.Name, "nonzero"); err != nil {
		log.Printf("Error to validate 'name': %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}
	if err := validator.Valid(data.CreatedBy, "nonzero"); err != nil {
		log.Printf("Error to validate 'created_by': %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}
	if err := validator.Valid(data.State, "nonzero"); err != nil {
		log.Printf("Error to validate 'state': %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}

	// 封装结构体
	t, err := models.AddTag(data.Name, data.State, data.CreatedBy)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.Error, data)
		return
	}
	data.ID = t.ID
	data.CreatedAt = t.CreatedAt
	data.UpdatedAt = t.UpdatedAt
	appG.Response(http.StatusCreated, e.Create, data)
}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}

func ExportTag(c *gin.Context) {

}

func ImportTag(c *gin.Context) {

}
