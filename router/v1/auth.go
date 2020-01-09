package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/model"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/util"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

// GetAuth 通过用户名、密码生成 token
// curl localhost:8080/api/v1/auth -H "username:admin" -H "password:123456"
func GetAuth(c *gin.Context) {
	var (
		username string
		password string
		code     int
		data     map[string]interface{}
	)

	username = c.GetHeader("username")
	password = c.GetHeader("password")
	code = e.InvalidParams

	if err := validator.Valid(username, "nonzero"); err != nil {
		log.Printf("Error to validate 'username': %v\n", err)
		c.JSON(http.StatusBadRequest, H(code, struct{}{}))
		return
	}

	if err := validator.Valid(password, "nonzero"); err != nil {
		log.Printf("Error to validate 'password': %v\n", err)
		c.JSON(http.StatusBadRequest, H(code, struct{}{}))
		return
	}

	b := model.CheckAuth(username, password)
	if !b {
		code = e.ErrorAuth
	} else {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ErrorAuthToken
		} else {
			code = e.Ok
			data["token"] = token
		}
	}

	c.JSON(http.StatusOK, H(code, data))
}
