# GrpcDebug
A grpc debugging tool with web ui

# 本地启动
```bash
# step1 创建db
cd sql
sqlite3 grpc_debug.db
# step2 启动程序
go mod download
cd cmd
go build -o ../main
```

# 目录介绍
/cmd    启动目录
/internal   内部实现
    -/dao   数据操作
    -/pkg   业务模块
    -/server    服务入口
    -/service   服务实现
/proto  存放grpc协议定义文件
/sql 数据库文件存放目录
/static 静态资源文件

# 常用命令
```bash
# 生成pb文件
cd proto
protoc --go_out=plugins=grpc:. hello.proto
```
