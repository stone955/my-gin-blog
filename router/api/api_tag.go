package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/model"
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

// GetTags 查询所有标签
// curl localhost:8080/api/tags
func GetTags(c *gin.Context) {
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

	data := make([]Tag, 0)

	tags, err := model.GetTags(util.GetPage(c), setting.PageSize, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, H(e.Error, data))
		return
	}

	if len(tags) > 0 {
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
	}

	c.JSON(http.StatusOK, H(e.Ok, data))
}

// AddTag 添加新标签
// curl -X POST localhost:8080/api/tags -d "{\"name\":\"golang\",\"state\":1,\"created_by\":\"admin\"}"
func AddTag(c *gin.Context) {
	var tag Tag
	if err := c.BindJSON(&tag); err != nil {
		log.Printf("Error to bind json: %v\n", err)
		c.JSON(http.StatusBadRequest, H(e.InvalidParams, struct{}{}))
		return
	}

	if err := validator.Valid(tag.Name, "nonzero"); err != nil {
		log.Printf("Error to validate 'name': %v\n", err)
		c.JSON(http.StatusBadRequest, H(e.InvalidParams, struct{}{}))
		return
	}
	if err := validator.Valid(tag.CreatedBy, "nonzero"); err != nil {
		log.Printf("Error to validate 'created_by': %v\n", err)
		c.JSON(http.StatusBadRequest, H(e.InvalidParams, struct{}{}))
		return
	}
	if err := validator.Valid(tag.State, "nonzero"); err != nil {
		log.Printf("Error to validate 'state': %v\n", err)
		c.JSON(http.StatusBadRequest, H(e.InvalidParams, struct{}{}))
		return
	}

	// 封装结构体
	t, err := model.AddTag(tag.Name, tag.State, tag.CreatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, H(e.Error, struct{}{}))
		return
	}
	tag.ID = t.ID
	tag.CreatedAt = t.CreatedAt
	tag.UpdatedAt = t.UpdatedAt
	c.JSON(http.StatusCreated, H(e.Create, &tag))
}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
