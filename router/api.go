package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/e"
)

func H(code int, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	}
}
