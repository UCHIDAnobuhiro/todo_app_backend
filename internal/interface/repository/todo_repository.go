package repository

import "todo_backend/internal/domain"

type TodoRepository interface {
	FindAll() ([]domain.Todo, error)
	Create(todo domain.Todo) error
}
