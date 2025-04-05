# User Service

CloudNative Days Handson用のユーザサービスのサンプルアプリケーションです。

---

## Features

- ユーザー登録
- ユーザーログイン
- アクセストークンの検証
- リフレッシュトークンを使用したアクセストークンの更新

---

## Requirements

- Go 1.24.1 以上
- PostgreSQL

---

## Setup

### 1. リポジトリをクローン
```bash
git clone https://github.com/cloudnativedaysjp/cnd-handson-app/
cd user-service
```

### 2. devContainerの環境変数を設定
`./devcontainr/.env` ファイルを作成し、以下のように設定してください（`.env.dummy` を参考にしてください）。

```env
DB_HOST=db
DB_PORT=5432
DB_DB=your_db_name
DB_USER=your_db_user
DB_PASSWORD=your_db_password
JWT_SECRET_KEY=your_jwt_secret_key
```

### 3. devcontainer起動
```bash
devcontainer open
```

### 4. マイグレーションを実行
```bash
go run cmd/server/main.go migrate
```

### 5. アプリケーションを起動
```bash
go run cmd/server/main.go serve
```
---
## AppのBuild
```bash
make
```
binにuser-serviceバイナリが作成されます。
## Quick Start

### ユーザー登録
```bash
curl -X POST http://localhost:8080/auth/register \
     -H "Content-Type: application/json" \
     -d '{
           "email": "test@example.com",
           "password": "securepassword"
         }'
```

### ユーザーログイン
```bash
curl -X POST http://localhost:8080/auth/login \
     -H "Content-Type: application/json" \
     -d '{
           "email": "test@example.com",
           "password": "securepassword"
         }'
```

### アクセストークンの検証
```bash
curl -X GET http://localhost:8080/auth/validate \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <your_access_token>"
```

### リフレッシュトークンを使用してアクセストークンを更新
```bash
curl -X POST http://localhost:8080/auth/refresh \
     -H "Content-Type: application/json" \
     -d '{
           "refresh_token": "<your_refresh_token>"
         }'
```

---

## API Documentation

APIの詳細な仕様は、[OpenAPI仕様書](openapi/openapi.yaml)を参照してください。Redocを使用してブラウザで確認することもできます。

### Redocを使用する場合
```bash
cd openapi
make build
make run
```
ブラウザで `http://localhost:8080` にアクセスしてください。

---

## Project Structure

```
.
├── cmd/                # エントリーポイント
│   └── server/
│       └── main.go     # メインアプリケーション
├── internal/           # 内部ロジック
│   ├── user/           # ユーザー関連
│   └── refresh/        # リフレッシュトークン関連
├── pkg/                # 再利用可能なパッケージ
│   ├── auth/           # 認証関連
│   └── db/             # データベース関連
├── openapi/            # OpenAPI仕様書
├── .devcontainer/      # DevContainer設定
├── Dockerfile          # Dockerビルド設定
├── docker-compose.yml  # Docker Compose設定
├── Makefile            # ビルド・テスト用Makefile
└── README.md           # このファイル
```

---



