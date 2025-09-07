package usecase

import (
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"
)

// TodoUsecase は、Todoエンティティに関するアプリケーション固有の
// ユースケースを実装する構造体です。
// Clean ArchitectureにおけるUsecase層であり、
// Repositoryインターフェースを通じて永続化層へアクセスします。
type TodoUsecase struct {
	Repo repository.TodoRepository
}

// NewTodoUsecaseは、指定されたTodoRepositoryを使用する
// TodoUsecaseの新しいインスタンスを返します。
func NewTodoUsecase(r repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{Repo: r}
}

// GetTodosは、登録されている全てのTodoを取得します。
func (uc *TodoUsecase) GetTodos(userID uint) ([]domain.Todo, error) {
	return uc.Repo.FindByUser(userID)
}

// AddTodoは、新しいTodoを作成して保存します。
func (uc *TodoUsecase) AddTodo(todo domain.Todo) error {
	return uc.Repo.Create(todo)
}

// UpdateTodoは、既存のTodoを更新します。
func (uc *TodoUsecase) UpdateTodo(todo domain.Todo) error {
	return uc.Repo.Update(todo)
}

// DeleteTodoは、指定されたIDのTodoを削除します。
func (uc *TodoUsecase) DeleteTodo(userID uint, id int) error {
	return uc.Repo.Delete(userID, id)
}
