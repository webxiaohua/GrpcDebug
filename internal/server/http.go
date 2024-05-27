/**
 * @author:伯约
 * @date:2024/5/26
 * @note:
**/

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/webxiaohua/GrpcDebug/internal/service"
	"net/http"
	"os"
	"strings"
)

type CommonResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewHttpServer() (engine *http.Server) {
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
	r.POST("/tool/grpc_debug/discovery_list", func(c *gin.Context) {
		var req service.GrpcDebugDiscoveryListReq
		if err := c.ShouldBindJSON(&req); err == nil {
			resp, _ := service.GrpcDebugDiscoveryList(c, &req)
			c.JSON(200, resp)
		} else {
			c.JSON(200, &CommonResp{
				Code:    400,
				Message: "参数错误",
				Data:    "",
			})
		}
	})
	r.POST("/tool/grpc_debug/node_list", func(c *gin.Context) {
		var req service.GrpcNodeListReq
		if err := c.ShouldBindJSON(&req); err == nil {
			resp, _ := service.GrpcNodeList(c, &req)
			c.JSON(200, resp)
		} else {
			c.JSON(200, &CommonResp{
				Code:    400,
				Message: "参数错误",
				Data:    "",
			})
		}
	})
	r.POST("/tool/grpc_debug", func(c *gin.Context) {
		var req service.GrpcDebugReq
		if err := c.ShouldBindJSON(&req); err == nil {
			resp, _ := service.GrpcDebug(c, &req)
			c.JSON(200, resp)
		} else {
			c.JSON(200, &CommonResp{
				Code:    400,
				Message: "参数错误",
				Data:    "",
			})
		}
	})
}

func getStaticPath() (staticPath string) {
	keyPathName := "GrpcDebug"
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
