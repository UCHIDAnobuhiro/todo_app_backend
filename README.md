todo_backend — Gin + Clean Architecture + GORM（日本語）

Go 製の最小 TODO バックエンド。HTTP は Gin、アーキテクチャは Clean Architecture、ORM は GORM を使用。ローカル開発は SQLite をデフォルトにし、DI で MySQL/PostgreSQL 等へ容易に差し替え可能です。

⸻

機能
• REST API: GET /todos, POST /todos, PUT /todos/:id, DELETE /todos/:id
• Clean Architecture（domain / usecase / interface / infrastructure）
• Repository インターフェース + 具体実装（DI で差し替え）
• GORM + AutoMigrate（初回起動で todo.db 作成）
• CORS 有効化（gin-contrib/cors）

⸻

技術スタック
• Go 1.24+
• Gin (github.com/gin-gonic/gin)
• GORM (gorm.io/gorm) + ドライバ（デフォルト SQLite）
• gin-contrib/cors

⸻

ディレクトリ構成（抜粋）

internal/
domain/ # エンティティ（ビジネスルール）
todo.go
usecase/ # アプリ固有のユースケース
todo_usecase.go
interface/
handler/ # Gin ハンドラ（HTTP 入出力）
todo_handler.go
repository/ # Repository の契約（インターフェース）
todo_repository.go
infrastructure/
mysql/ # Repository 実装（GORM 利用）
todo_mysql.go
main.go # 組み立て（DI）・サーバ起動

⸻

セットアップ

1. 依存解決

go mod tidy

2. 起動（SQLite デフォルト）

go run ./...

# サーバ: http://localhost:8080

起動時に以下が実行されます：
• todo.db（SQLite）の作成/接続
• domain.Todo を AutoMigrate でテーブル同期
• Gin（CORS 有効）で HTTP サーバ起動

⸻

API 仕様（現状の実装準拠）

GET /todos
• 200: [{ id, title, completed }, ...]

POST /todos
• リクエスト Body（JSON）: { "title": string, "completed": bool }
• 201: { "message": "created" }

PUT /todos/:id
• パスパラメータ: :id
• リクエスト Body（JSON）: { id, title, completed }（実装上、全体を送る想定）
• 200: { "message": "updated" }

DELETE /todos/:id
• 200: { "message": "deleted" }

必要に応じて、リクエスト/レスポンスのスキーマやエラーレスポンス仕様を明確化してください（例：バリデーションエラー時の 400 フォーマットなど）。

⸻

エンティティ（domain）

// internal/domain/todo.go
type Todo struct {
ID uint `json:"id"`
Title string `json:"title"`
Completed bool `json:"completed"`
}

⸻

Repository（契約と実装）

契約（interface）

// internal/interface/repository/todo_repository.go
type TodoRepository interface {
FindAll() ([]domain.Todo, error)
Create(todo domain.Todo) error
Update(todo domain.Todo) error
Delete(id int) error
}

実装（GORM + SQLite/MySQL 等）

// internal/infrastructure/mysql/todo_mysql.go
type TodoMysql struct { DB *gorm.DB }
func (r *TodoMysql) FindAll() ([]domain.Todo, error) { var t []domain.Todo; return t, r.DB.Find(&t).Error }
func (r *TodoMysql) Create(todo domain.Todo) error { return r.DB.Create(&todo).Error }
func (r *TodoMysql) Update(todo domain.Todo) error { return r.DB.Save(&todo).Error }
func (r \*TodoMysql) Delete(id int) error { return r.DB.Delete(&domain.Todo{}, id).Error }

⸻

Usecase（ユースケース）

// internal/usecase/todo_usecase.go
type TodoUsecase struct { Repo repository.TodoRepository }
func NewTodoUsecase(r repository.TodoRepository) *TodoUsecase { return &TodoUsecase{Repo: r} }
func (uc *TodoUsecase) GetTodos() ([]domain.Todo, error) { return uc.Repo.FindAll() }
func (uc *TodoUsecase) AddTodo(todo domain.Todo) error { return uc.Repo.Create(todo) }
func (uc *TodoUsecase) UpdateTodo(todo domain.Todo) error{ return uc.Repo.Update(todo) }
func (uc \*TodoUsecase) DeleteTodo(id int) error { return uc.Repo.Delete(id) }

⸻

Handler（Gin）

// internal/interface/handler/todo_handler.go
func (h *TodoHandler) GetTodos(c *gin.Context) {
todos, err := h.Usecase.GetTodos()
if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
c.JSON(http.StatusOK, todos)
}

⸻

main.go（組み立て + 起動）

// DB（SQLite）
db, \_ := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
db.AutoMigrate(&domain.Todo{})

// DI（repo → usecase → handler）
repo := mysql.NewTodoMysql(db)
uc := usecase.NewTodoUsecase(repo)
r := gin.Default()
r.Use(cors.Default())
handler.NewTodoHandler(r, uc)
r.Run(":8080")

DB を差し替える（例：MySQL）

import "gorm.io/driver/mysql"
dsn := "user:pass@tcp(localhost:3306)/todo?parseTime=true&charset=utf8mb4"
db, \_ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
repo := mysql.NewTodoMysql(db) // 他はそのまま

⸻

curl での確認例

# 一覧

curl -i http://localhost:8080/todos

# 追加

curl -i -X POST http://localhost:8080/todos \
 -H 'Content-Type: application/json' \
 -d '{"title":"Buy milk","completed":false}'

# 更新

curl -i -X PUT http://localhost:8080/todos/1 \
 -H 'Content-Type: application/json' \
 -d '{"id":1,"title":"Buy milk & bread","completed":true}'

# 削除

curl -i -X DELETE http://localhost:8080/todos/1

⸻

Clean Architecture と DI の関係
• domain: エンティティ・不変条件（技術非依存）
• usecase: アプリの手順（Repository の契約にのみ依存）
• interface/handler: HTTP の入出力（Gin）
• interface/repository: データ操作の契約
• infrastructure/mysql: 契約の実装（GORM/DB）
• main.go: 実装選択と**依存性注入（DI）**の場所（差し替えポイント）

依存は外側 → 内側。DB 変更時は main.go の注入箇所のみ 差し替えれば OK。

⸻

開発メモ
• CORS: r.Use(cors.Default()) を必要に応じてカスタマイズ
• ドキュメント: godoc -http=:6060 でコメントから HTML ドキュメント生成
• フォーマット: gofmt, goimports を使用
• 設定: 将来的には環境変数で DSN/ポートを管理するのが ◎
• エラーハンドリング: 現状 500 で丸めて返却。要件に応じて 400/404 等の詳細化推奨

⸻

トラブルシュート
• DB ファイルができない: 実行ディレクトリ/権限を確認
• マイグレーションされない: AutoMigrate の呼び出し順序を確認
• CORS でブロック: 許可オリジン/ヘッダ/メソッドを cors.Config で明示
• ポート競合: r.Run(":8080") のポート番号を変更

⸻
