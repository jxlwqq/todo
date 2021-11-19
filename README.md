# Todo 服务

### 环境依赖

* docker-desktop >= 4.2.0
* kubernetes >= 1.21.5
* go >= 1.17
* istioctl >= 1.11.4
* protobuf >= 3.17.3
* grpcurl >= 1.8.5


下载安装 Docker Desktop ，并启动内置的 Kubernetes 集群。

```shell
# 安装 Go
brew install go
# 安装 Protobuf
brew install protobuf
# 安装 grpcurl
brew install grpcurl
# 安装 Istio
brew install istioctl
kubectl config use-context docker-desktop
istioctl install -y
kubectl label namespace default istio-injection=enabled
```

### Makefile 介绍

|  命令   | 说明  |
|  ----  | ----  |
| `make init`  | 安装 protoc-gen-go、protoc-gen-grpc 和 wire|
| `make protoc`  | 基于 proto 文件，生成 *_pb.go 和 *_grpc.pb.go |
| `make wire`    | 基于 wire.go文件，生成 wire_gen.go |
| `make docker-build`  | 构建 docker 镜像 |
| `make docker-push`   | 推送 docker 镜像 |
| `make kube-deploy-mysql` | 在集群中部署 MySQL 服务 |
| `make kube-deploy-todo` | 在集群中部署 Todo 服务 |
| `make kube-deploy-istio` | 在集群中部署 Istio Gateway 和 VirtualService |
| `make kube-deploy-all` | 在集群中部署所有资源 |
| `make kube-delete-all` | 在集群中删除所有资源 |

### 构建镜像

```shell
make docker-build
```

### 部署服务

部署 MySQL 服务：

```shell
make kube-deploy-mysql
```

部署 Todo 服务：

```shell
make kube-deploy-todo
```

部署 Istio Gateway 和 VirtualService：

```shell
make kube-deploy-istio
```

### 使用 grpcurl 访问 gRPC 服务

```shell
grpcurl -plaintext 127.0.0.1:80 list
```

返回：

```shell
grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection
v1.Todo
```

CRUD 操作：

```shell
# create
grpcurl -d '{"item": {"title":"10点会议", "description": "服务架构优化", "remind_at": "2021-01-02T15:04:05.999999999Z"}}' -plaintext 127.0.0.1:80  v1.Todo.Create 
# get
grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:80  v1.Todo.Get 
# list
grpcurl -plaintext 127.0.0.1:80  v1.Todo.List
# update
grpcurl -d '{"item": {"id": 1, "title":"10点会议", "description": "服务架构调整", "remind_at": "2021-01-02T15:04:05.999999999Z"}}' -plaintext 127.0.0.1:80  v1.Todo.Update
# delete
grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:80  v1.Todo.Delete 
```