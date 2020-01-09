package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/stone955/my-gin-blog/pkg/setting"
	"github.com/stone955/my-gin-blog/router"
	"log"
	"syscall"
)

func main() {
	r := router.Register()

	/*
		// 优雅停机：自己实现
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
		signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
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
	*/

	// 优雅停机：使用 endless 实现
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d\n", syscall.Getpid())
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server err: %v\n", err)
	}
}
