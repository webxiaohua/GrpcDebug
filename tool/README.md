# 使用说明
```bash
cd tool
go install grpcDebugCli
gprcDebugCli --addr=127.0.0.1:9000 --data='{"name":"Timi","ids":[1,2,3]}' --path=/pb.Greeter/SayHello
```