package main

import (
	"context"
	"fmt"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	r := Register()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:      r,
		ReadTimeout:  setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
	}

	var wg sync.WaitGroup
	// 优雅关闭
	exit := make(chan os.Signal)
	//监听 Ctrl+C 信号
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-exit
		wg.Add(1)
		defer wg.Done()
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
}
