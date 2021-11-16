# Todo 服务

### 使用 grpcurl 访问gRPC服务

```shell
brew install grpcurl
```

```shell
# create
grpcurl -d '{"todo": {"title":"123", "description": "123", "remind_at": "2006-01-02T15:04:05.999999999Z"}}' -plaintext 127.0.0.1:50051  v1.TodoService.Create 
# get
grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:50051  v1.TodoService.Get 
# list
grpcurl -plaintext 127.0.0.1:50051  v1.TodoService.List
# update
grpcurl -d '{"todo": {"id": 1, "title":"123456", "description": "123456", "remind_at": "2006-01-02T15:04:05.999999999Z"}}' -plaintext 127.0.0.1:50051  v1.TodoService.Update
# delete
grpcurl -d '{"id": 1}' -plaintext 127.0.0.1:50051  v1.TodoService.Delete 
```