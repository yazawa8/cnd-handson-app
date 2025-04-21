all: bin/session-service
test: lint unit-test
IMAGE_NAME=session-service
TAG=latest
PLATFORM?=linux/amd64

.PHONY: bin/session-service
bin/session-service:
	@docker build . --target bin -t $(IMAGE_NAME):$(TAG) \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: bin/session-client
bin/session-client:
	@docker build . --target client -t $(IMAGE_NAME):$(TAG) \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test \
	$(BUILD_DIR)

.PHONY: unit-test-coverage
unit-test-coverage:
	@docker build . --target unit-test-coverage \
	$(BUILD_DIR) \
	--output coverage/ 
	cat coverage/cover.out

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
