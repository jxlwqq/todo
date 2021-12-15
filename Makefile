.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: protoc
protoc:
	for file in $$(git ls-files '*.proto'); do \
		protoc -I $$(dirname $$file) \
		--go_out=:$$(dirname $$file) --go_opt=paths=source_relative \
		--go-grpc_out=:$$(dirname $$file) --go-grpc_opt=paths=source_relative \
		$$file; \
	done

.PHONY: wire
wire:
	wire ./cmd/server

.PHONY: docker-build
docker-build:
	docker build -t jxlwqq/todo:latest -f ./build/docker/server/Dockerfile .

.PHONY: docker-push
docker-push:
	docker push jxlwqq/todo:latest

.PHONY: docker-buildx
docker-buildx:
	docker buildx build -t jxlwqq/todo:latest --platform linux/arm64,linux/amd64 --push -f ./build/docker/server/Dockerfile .

.PHONY: kube-deploy-mysql
kube-deploy-mysql:
	kubectl apply -f deployments/mysql.yaml
	kubectl rollout status deployments/mysql

.PHONY: kube-deploy-todo
kube-deploy-todo:
	kubectl apply -f deployments/todo.yaml
	kubectl rollout status deployments/todo

.PHONY: kube-deploy-istio
kube-deploy-istio:
	kubectl apply -f deployments/istio.yaml

.PHONY: kube-deploy-all
kube-deploy-all:
	make kube-deploy-mysql
	make kube-deploy-todo
	make kube-deploy-istio

.PHONY: kube-delete-all
kube-delete-all:
	kubectl delete -f deployments/