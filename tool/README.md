# 使用说明
```bash
# 进入grpcDebugCli目录
cd tool/grpcDebugCli
# 安装客户端
go install
# 运行客户端
gprcDebugCli --addr=127.0.0.1:9000 --data='{"name":"Timi","ids":[1,2,3]}' --path=/pb.Greeter/SayHello
```