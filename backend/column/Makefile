all: bin/user-service
test: lint unit-test
IMAGE_NAME=user-service
TAG=latest
PLATFORM?=linux/amd64

.PHONY: bin/user-service
runtime:
	@docker build . --target runtime -t $(IMAGE_NAME):$(TAG)
.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: proto
proto: 
	@protoc --proto_path=./proto --proto_path=./proto \
	--go_out=./proto --go-grpc_out=./proto proto/user.proto
	@mv proto/github.com/cloudnativedaysjp/cnd-handson-app/backend/user/proto/* proto/
	@rm -rf proto/github.com
	@protoc --proto_path=./proto \
        --doc_out=./proto --doc_opt=markdown,README.md \
        proto/user.proto
	@protoc --proto_path=./proto \
        --doc_out=./proto --doc_opt=html,index.html \
        proto/user.proto
