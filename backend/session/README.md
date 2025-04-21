# Session Service

CloudNative Days Handson用のsessionサービスのサンプルアプリケーションです。

---

## Features

- ユーザーログイン
- アクセストークンの発行
- アクセストークンの検証
- リフレッシュトークンの発行
- リフレッシュトークンの検証
- リフレッシュトークンの無効化（logout）
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
cd session-service
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
binにsession-serviceバイナリが作成されます。
## Quick Start

---


---

## Project Structure

```
.
├── cmd/                # エントリーポイント
│   └── server/
│       └── main.go     # メインアプリケーション
├── internal/           # 内部ロジック
│   ├── session/           # ユーザー関連
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



