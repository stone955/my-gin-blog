package main

import (
	"context"
	"fmt"
	"github.com/stone955/my-gin-blog/models"
	"github.com/stone955/my-gin-blog/pkg/gredis"
	"github.com/stone955/my-gin-blog/pkg/logging"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	setting.Setup()
	logging.Setup()
	gredis.Setup()
	models.Setup()

	r := router.Register()

	// 优雅停机：自己实现
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.ServerCfg.HttpPort),
		Handler:      r,
		ReadTimeout:  setting.ServerCfg.ReadTimeout,
		WriteTimeout: setting.ServerCfg.WriteTimeout,
	}

	var wg sync.WaitGroup
	// 优雅关闭
	exit := make(chan os.Signal)
	//监听 Ctrl+C 信号
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-exit
		wg.Add(1)
		defer wg.Done()
		// 关闭数据库连接
		defer models.Close()
		// 关闭缓存连接
		defer gredis.Close()
		//使用context控制srv.Shutdown的超时时间
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Fatal to shutdown: %v\n", err)
		}
	}()

	// 启动 server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Fatal to listen and serve: %v\n", err)
	}

	wg.Wait()

	log.Println("Gracefully shutdown the http server...")

	/*
		// 优雅停机：使用 endless 实现
		endless.DefaultReadTimeOut = setting.ServerCfg.ReadTimeout
		endless.DefaultWriteTimeOut = setting.ServerCfg.WriteTimeout
		endless.DefaultMaxHeaderBytes = 1 << 20
		endPoint := fmt.Sprintf(":%d", setting.ServerCfg.HTTPPort)

		server := endless.NewServer(endPoint, r)
		server.BeforeBegin = func(add string) {
			log.Printf("Actual pid is %d\n", syscall.Getpid())
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Server err: %v\n", err)
		}
	*/

}
