## todo_backend — Gin + クリーンアーキテクチャ + GORM

Go 言語で構築したシンプルな TODO バックエンドです。Gin（HTTP サーバ）、クリーンアーキテクチャ（domain / usecase / interface / infrastructure）、GORM（ORM）を採用しています。デフォルトでは SQLite を利用しますが、DI により MySQL / PostgreSQL などに容易に差し替え可能です。

---

### 特徴

- RESTful API: GET/POST/PUT/DELETE /todos
- クリーンアーキテクチャ構成 + 依存性注入（Repository インターフェースと実装の分離）
- GORM + AutoMigrate によるスキーマ自動生成
- CORS 設定済み（gin-contrib/cors）

---

### 技術スタック

- Go 1.24+
- Gin (github.com/gin-gonic/gin)
- GORM (gorm.io/gorm) + ドライバ（SQLite デフォルト）
- gin-contrib/cors

---

### ディレクトリ構成（簡略）

```
internal/
    domain/ # エンティティ（純粋なビジネスルール）
    usecase/ # アプリケーション固有のユースケース
    interface/
        handler/ # Gin ハンドラー（HTTP I/O）
        repository/ # Repository インターフェース（契約定義）
    infrastructure/
        mysql/ # Repository 実装（GORM 使用）
main.go # 各層の接続とサーバ起動（Composition Root）
```

---

### 実行方法

1. クローン & 依存関係取得

```
git clone <repo-url>
cd todo_backend
go mod tidy
```

2. 起動（SQLite デフォルト）

```
go run cmd/main.go
```

- todo.db が作成され、domain.Todo のスキーマが自動適用されます
- Gin サーバが http://localhost:8080 で起動します

3. 簡単な動作確認

```
# 全件取得
curl -i http://localhost:8080/todos

# 作成
curl -i -X POST http://localhost:8080/todos \
 -H 'Content-Type: application/json' \
 -d '{"title":"牛乳を買う","completed":false}'

# 更新
curl -i -X PUT http://localhost:8080/todos/1 \
 -H 'Content-Type: application/json' \
 -d '{"id":1,"title":"牛乳とパンを買う","completed":true}'

# 削除
curl -i -X DELETE http://localhost:8080/todos/1
```

---

### API 仕様

- GET /todos → 登録済み TODO 一覧取得
- POST /todos → 新規作成
- PUT /todos/:id → 更新
- DELETE /todos/:id → 削除

レスポンス例:

```
[
{"id":1, "title":"牛乳を買う", "completed":false}
]
```

---

### クリーンアーキテクチャと DI

- domain: エンティティ（例: Todo）とビジネスルール
- usecase: ユースケースの流れ（TodoUsecase）
- interface/repository: DB アクセス契約（TodoRepository）
- infrastructure/mysql: 実際の DB 実装（GORM）
- interface/handler: Gin ハンドラー（HTTP リクエストとユースケースを接続）
- main.go: 実装を選択して注入し、サーバを起動

💡 DB を差し替える場合は main.go の接続部分だけを変更すれば OK です。
