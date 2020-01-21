package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/models"
	"github.com/stone955/my-gin-blog/pkg/app"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/util"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

// GetAuth 通过用户名、密码生成 token
// curl 192.168.1.108:8080/auth -H "username:admin" -H "password:123456"
func GetAuth(c *gin.Context) {
	var (
		appG     = app.Gin{C: c}
		username = c.GetHeader("username")
		password = c.GetHeader("password")
		data     = make(map[string]interface{})
	)

	log.Printf("Request header username: %v, password: %v\n", username, password)

	if err := validator.Valid(username, "nonzero"); err != nil {
		log.Printf("Error to validate 'username': %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}

	if err := validator.Valid(password, "nonzero"); err != nil {
		log.Printf("Error to validate 'password': %v\n", err)
		appG.Response(http.StatusBadRequest, e.InvalidParams, data)
		return
	}

	if !models.CheckAuth(username, password) {
		appG.Response(http.StatusUnauthorized, e.ErrorAuth, data)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ErrorAuthToken, data)
		return
	}

	data["token"] = token
	appG.Response(http.StatusOK, e.Ok, data)
}
