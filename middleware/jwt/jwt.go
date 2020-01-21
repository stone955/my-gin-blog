package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/stone955/my-gin-blog/pkg/app"
	"github.com/stone955/my-gin-blog/pkg/e"
	"github.com/stone955/my-gin-blog/pkg/util"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		var (
			appG = app.Gin{C: c}
			code int
			data interface{}
		)

		code = e.Ok
		token := c.Query("token")
		if token == "" {
			code = e.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}

		if code != e.Ok {
			appG.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}
	}
}
