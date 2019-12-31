package main

import (
	"fmt"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/router"
	"net/http"
)

func main() {
	r := router.Register()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:      r,
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
	}
	s.ListenAndServe()
}
