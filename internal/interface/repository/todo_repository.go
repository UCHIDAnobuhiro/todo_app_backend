package repository

import "todo_backend/internal/domain"

type TodoRepository interface {
	FindAll() ([]domain.Todo, error)
	Create(todo domain.Todo) error
	Update(todo domain.Todo) error
	Delete(id int) error
}
