/**
 * @author:伯约
 * @date:2024/5/26
 * @note:
**/

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func New() (engine *http.Server) {
	fmt.Println("hello")
	r := gin.Default()
	router(r)
	// 启动 HTTP 服务器
	engine = &http.Server{
		Addr:    ":8000",
		Handler: r,
	}
	return
}

func router(r *gin.Engine) {
	staticPath := getStaticPath() //自适应获取目录地址
	r.Static("/static", staticPath)
	r.GET("/version", func(c *gin.Context) {
		c.JSON(200, "1.0.0")
	})
}

func getStaticPath() (staticPath string) {
	keyPathName := "GrpcDebug"
	defer func() {
		_ = os.Setenv("STATIC_PATH", staticPath)
	}()
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("无法获取当前工作目录：%s\n", err))
	}
	if strings.Contains(wd, keyPathName) {
		staticPath = "./static"
		return
	}
	staticPath = "./GrpcDebug/static"

	return
}
