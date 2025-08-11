package mysql

import (
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"

	"gorm.io/gorm"
)

// TodoMysqlはGORMを利用してMySQLデータベース上で
// Todoエンティティの永続化処理を行う構造体です。
// CleanArchitectureにおけるInfrastructure層に相当します。
type TodoMysql struct {
	DB *gorm.DB
}

// コンパイル時に TodoMysql が repository.TodoRepository インターフェースを
// 実装しているか確認します（実装漏れ防止）。
var _ repository.TodoRepository = (*TodoMysql)(nil)

// NewTodoMysql は、指定された gorm.DB 接続を使用する TodoMysqlの
// 新しいインスタンスを返します（DI用のコンストラクタ）。
func NewTodoMysql(db *gorm.DB) *TodoMysql {
	return &TodoMysql{DB: db}
}

// FindAll は、データベースから全ての Todo を取得します。
func (r *TodoMysql) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.DB.Find(&todos).Error
	return todos, err
}

// Create は、指定されたTodoをデータベースに新規登録します。
func (r *TodoMysql) Create(todo domain.Todo) error {
	return r.DB.Create(&todo).Error
}

// Update は、指定されたTodoの情報をデータベース上で更新します。
func (r *TodoMysql) Update(todo domain.Todo) error {
	return r.DB.Save(&todo).Error
}

// Delete は、指定されたIDのTodoをデータベースから削除します。
func (r *TodoMysql) Delete(id int) error {
	return r.DB.Delete(&domain.Todo{}, id).Error
}
