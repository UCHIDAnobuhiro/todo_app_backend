package mysql

import (
	"todo_backend/internal/domain"
	"todo_backend/internal/interface/repository"

	"gorm.io/gorm"
)

type TodoMysql struct {
	DB *gorm.DB
}

var _ repository.TodoRepository = (*TodoMysql)(nil)

func NewTodoMysql(db *gorm.DB) *TodoMysql {
	return &TodoMysql{DB: db}
}

func (r *TodoMysql) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.DB.Find(&todos).Error
	return todos, err
}

func (r *TodoMysql) Create(todo domain.Todo) error {
	return r.DB.Create(&todo).Error
}
