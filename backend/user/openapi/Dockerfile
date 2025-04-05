# ビルドに必要なベースイメージ
FROM node:16 AS build

WORKDIR /app

# 必要なファイルをコピー
COPY ./ /app/

# Redoc CLIをインストール
RUN npm install -g redoc-cli

# OpenAPI仕様書をHTMLに変換
RUN redoc-cli bundle openapi.yaml --output /app/index.html

# 最終的に生成された静的ファイルをホスト用に配置する
FROM nginx:alpine

# Nginxのデフォルトの公開ディレクトリにコピー
COPY --from=build /app/index.html /usr/share/nginx/html/index.html

# Nginxを公開ポートで実行
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
