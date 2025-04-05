# イメージ名とタグの定義
IMAGE_NAME=openapi-server
TAG=latest


# Dockerビルド
build:
	docker build -t $(IMAGE_NAME):$(TAG) .

# コンテナ起動
run:
	docker run -d -p 8080:80 --name $(IMAGE_NAME) $(IMAGE_NAME):$(TAG)

# コンテナ停止
stop:
	docker stop $(IMAGE_NAME) && docker rm $(IMAGE_NAME)

# log
logs:
	docker logs $(IMAGE_NAME)
