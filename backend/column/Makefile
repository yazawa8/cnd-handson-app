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
	@python -m grpc_tools.protoc -I. --python_out=./proto --grpc_python_out=./proto ./proto/column.proto
	@mv ./proto/proto/* ./proto/
	@rm -rf ./proto/proto

