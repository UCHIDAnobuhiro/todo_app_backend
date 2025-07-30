package usecase

import (
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"
)

type TodoUsecase struct {
	Repo repository.TodoRepository
}

func NewTodoUsecase(r repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{Repo: r}
}

func (uc *TodoUsecase) GetTodos() ([]domain.Todo, error) {
	return uc.Repo.FindAll()
}

func (uc *TodoUsecase) AddTodo(todo domain.Todo) error {
	return uc.Repo.Create(todo)
}
