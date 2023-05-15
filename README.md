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