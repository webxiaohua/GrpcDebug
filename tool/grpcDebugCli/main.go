/**
 * @author:伯约
 * @date:2024/5/28
 * @note: grpcDebug 命令行工具
**/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/webxiaohua/GrpcDebug/internal/pkg"
	"os"
)

var addr, method, data string
var RootCommand = &cobra.Command{
	Use:   "grpcDebugCli",
	Short: "grpcDebugCli is a tool for debug grpc service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addr:", addr)
		fmt.Println("method:", method)
		fmt.Println("data:", data)
		reply, err := pkg.GrpcInvoke(context.TODO(), addr, method, data)
		if err != nil {
			fmt.Println("grpcInvoke error:", err)
		} else {
			replyStr, _ := json.Marshal(reply)
			fmt.Printf("grpcInvoke success: \n\n%s\n\n", replyStr)
		}
	},
}

func init() {
	/*flag.StringVar(&data, "data", `{}`, `grpc service request data with json format`)
	flag.StringVar(&addr, "addr", `127.0.0.1:9000`, `grpc service address`)
	flag.StringVar(&method, "path", "/testproto.Greeter/SayHello", `grpc service path`)*/

	cobra.OnInitialize()
	RootCommand.Flags().StringVarP(&addr, "addr", "a", "127.0.0.1:9000", "grpc service address")
	RootCommand.Flags().StringVarP(&method, "path", "p", "/", "grpc service path")
	RootCommand.Flags().StringVarP(&data, "data", "d", "{}", "grpc service request data with json format")
}

func main() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
