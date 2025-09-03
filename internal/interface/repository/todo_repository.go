package repository

import "todo_backend/internal/domain"

// TodoRepository は Todo エンティティの永続化操作を定義するインターフェースです。
// Clean Architecture における Repository 層の契約を表し、
// 実際のデータストア（MySQL、PostgreSQL、メモリなど）の実装はこのインターフェースを満たす必要があります。
type TodoRepository interface {
	// FindAll は、指定ユーザーの全ての Todo を取得します。
	// 戻り値は Todo のスライスと、エラー情報です。
	FindByUser(userId uint) ([]domain.Todo, error)

	// Create は、新しい Todo を永続化します。
	// 引数には作成する Todo エンティティを渡します。
	Create(todo domain.Todo) error

	// Update は、既存の Todo を更新します。
	// 引数には更新内容を含む Todo エンティティを渡します。
	Update(todo domain.Todo) error

	// Delete は、指定された ID の Todo を削除します。
	Delete(userID uint, id int) error
}
