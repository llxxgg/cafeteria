# cafeteria

Using gin and grpc-go framework to build a cafeteria project;

Include Microsoft service module: gorm/metrics/trace/log/docker compose/k8s/DDD

Start from 2023.11.10

## Reference

[Gin框架中文文档](https://learnku.com/docs/gin-gonic/1.5/quickstart/6151)

## Todolist

* 实现remote rpc连接池
* telemetry
* metrics
* metadata使用
* protobuf使用 ‼️
* Grpc 如何保持连接的，keepalive？

## protobuf

```shell
grpcurl -plaintext localhost:8090 list

grpcurl -plaintext -d '{"name": "chris"}' localhost:8090 proto.Greeter.SayHello
```
使用grpcurl调试，前提需要grpc server注册了反射服务。

-plaintext grpcurl工具默认使用tls认证，如果服务是非tls认证的，可以通过指定这个选项来忽略tls认证。

protoc-gen-go-gin: proto gin http server

struct tag生成: https://github.com/favadi/protoc-go-inject-tag

## makefile




