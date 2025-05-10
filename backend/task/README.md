# Task Service

CloudNative Days Handson用のタスクサービスのサンプルアプリケーションです。

---

## Features

- タスク作成
- タスク更新
- タスク情報取得
- タスク削除
- タスク一覧

---

## Requirements

- Go 1.24.1 以上
- PostgreSQL

---

## Setup

### 1. リポジトリをクローン
```bash
git clone https://github.com/cloudnativedaysjp/cnd-handson-app/
cd cnd-handson-app/backend/task
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
もし失敗するorテーブルをリセットしたい場合
```bash
go run cmd/server/main.go reset
```

### 5. アプリケーションを起動
```bash
go run cmd/server/main.go server
```
---
## AppのBuild
```bash
make
```
binにtask-serviceバイナリが作成されます。
## Quick Start

---

## gRPC Documentation

gRPCの詳細な仕様は、[gRPC仕様書](proto/READEME.md)を参照してください。


## Project Structure

```
.
├── cmd/                # エントリーポイント
│   ├── client/         # clientのエントリーポイント
│   └── server/         # serverのエントリーポイント
├── internal/           # 内部ロジック
│   └── user/           # ユーザー関連
├── pkg/                # 再利用可能なパッケージ
│   ├── auth/           # 認証関連
│   └── db/             # データベース関連
├── proto/              # proto関連
├── .devcontainer/      # DevContainer設定
├── Dockerfile          # Dockerビルド設定
├── docker-compose.yml  # Docker Compose設定
├── Makefile            # ビルド・テスト用Makefile
└── README.md           # このファイル
```

---



