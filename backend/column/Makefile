all: bin/column-service
test: lint unit-test
IMAGE_NAME=column-service
TAG=latest
PLATFORM?=linux/amd64

.PHONY: bin/column-service
runtime:
	@docker build . --target runtime -t $(IMAGE_NAME):$(TAG)
.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: proto
proto: 
	@protoc --proto_path=./proto --proto_path=./proto \
	--go_out=./proto --go-grpc_out=./proto proto/column.proto
	@mv proto/github.com/cloudnativedaysjp/cnd-handson-app/backend/column/proto/* proto/
	@rm -rf proto/github.com
	@protoc --proto_path=./proto \
        --doc_out=./proto --doc_opt=markdown,README.md \
        proto/column.proto
	@protoc --proto_path=./proto \
        --doc_out=./proto --doc_opt=html,index.html \
        proto/column.proto
