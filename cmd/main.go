/**
 * @author:伯约
 * @date:2024/5/26
 * @note:
**/

package main

import (
	"context"
	"fmt"
	"github.com/webxiaohua/GrpcDebug/internal/server"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	httpServer := server.NewHttpServer()
	grpcServer := server.NewGrpcServer()
	// 使用 WaitGroup 来等待所有请求处理完成
	var wg sync.WaitGroup
	// 启动 HTTP 服务器的 goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
	}()
	// 监听停机信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// 等待停机信号
	<-signalChan
	fmt.Println("Received shutdown signal. Initiating graceful shutdown...")

	// 使用 context 来设置停机超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 停止 HTTP 服务器，并等待所有请求处理完成
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error during http server shutdown: %v\n", err)
	}
	if err := grpcServer.Stop; err != nil {
		fmt.Printf("Error during grpc server shutdown: %v\n", err)
	}
	// 等待所有请求处理完成
	wg.Wait()
	fmt.Println("Graceful shutdown complete.")
}
