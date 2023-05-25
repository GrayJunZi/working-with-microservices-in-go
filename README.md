# Anxiety

Working with Microservices in Go
---
Build highly available, scalable, resilient distributed applications using Go

## 一、介绍

- 微服务 (Microservices)
    1. 将整体从 functions/packages 分解为完全独立的程序。
    2. 通过 JSON/REST、RPC、gRPC和消息队列进行通信。
    3. 更容易扩展。
    4. 跟容易维护。

- 项目内容
    1. `Broker` - 可选的微服务入口。
    2. `Authentication` - Postgres, 身份认证服务。
    3. `Logger` - MongoDB, 日志服务。
    4. `Mail` - 发送电子邮件与特定模板。
    5. `Listener` - 消费RabbitMQ中的消息并启动一个进程。
- 通过以下方式在微服务之间进行通信。
    1. 使用JSON进行传输的 REST API.
    2. 使用RPC和gRPC进行发送和接收。
    4. 使用高级消息队列协议(AMQP)启动和响应时间。

一个前端web应用程序连接5个微服务。

## 二、代理服务(Broker Service)

安装依赖
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

## 三、认证服务(Authentication Service)

初始化
```bash
go init authenticaiton
go work use authentication-service/
go mod tidy
```

安装依赖
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/jackc/pgx/stdlib
go get -u github.com/tsawler/toolbox
```

## 四、日志服务(Logger Service)

初始化
```bash
go mod init logger
go work use logger-service
go mod tidy
```

安装依赖
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
```

## 五、邮件服务(Mail Service)

初始化
```bash
go mod init mail
go work use mail-service
```

安装依赖
```bash
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/go-chi/cors
go get github.com/vanng822/go-premailer/premailer
go get github.com/xhit/go-simple-mail/v2
```

## 六、监听服务(Listener Service)

初始化
```bash
go mod init listener
go work use listener-service
```

安装依赖
```bash
go get github.com/rabbitmq/amqp091-go
```

`broker-service` 安装依赖
```bash
go get github.com/rabbitmq/amqp091-go
```

## 七、RPC

## 八、gRPC

安装gRPC工具
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

生成gRPC代码
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative logs.proto
```

安装依赖
```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
```

## 九、Docker Swarm容器编排 

构建镜像
```bash
docker build -f logger-service.dockerfile -t grayjunzi/logger-service:1.0.0 .
```

推送镜像到 Docker Hub 中
```bash
docker push grayjunzi/logger-service:1.0.0
```

登录 docker
```bash
docker login
```

初始化docker swarm
```bash
docker swarm init
```

查看 加入worker命令 
```bash
docker swarm join-token worker
```

查看 加入manager命令 
```bash
docker swarm join-token manager
```

部署swarm
```bash
docker stack deploy -c swarm.yml myapp
```


查看所有服务
```bash
docker service ls
```

### 伸缩服务(Scaling services)

伸缩服务为3个实例
```bash
docker service scale myapp_listener-service=3
```

### 更新服务(Updating services)

构建新版本镜像
```bash
docker build -f logger-service.dockerfile -t grayjunzi/logger-service:1.0.1 .
```

推送镜像到docker hub
```bash
docker push grayjunzi/logger-service:1.0.1
```

扩展服务规模
```bash
dcoker service scale myapp_logger-service=2
```

更新服务
```bash
docker service update --image grayjunzi/logger-service:1.0.1 myapp_logger-service
```

查看服务
```bash
dcoker service ls
```

### 停止 Docker Swarm

将服务实例设置为0
```bash
docker service scale myapp_logger-service=0
```

移除所有服务
```bash
docker stack rm myapp
```

移除swarm
```bash
docker swarm leave
```

强制移除swarm
```bash
docker swarm leave -f
```

### Caddy

Caddy 是一个开源Web服务器