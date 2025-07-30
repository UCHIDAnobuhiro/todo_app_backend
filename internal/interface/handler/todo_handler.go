package handler

import (
	"net/http"
	"todo_backend/internal/domain"
	"todo_backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Usecase *usecase.TodoUsecase
}

func NewTodoHandler(r *gin.Engine, uc *usecase.TodoUsecase) {
	h := &TodoHandler{Usecase: uc}
	r.GET("/todos", h.GetTodos)
	r.POST("/todos", h.CreateTodo)
}

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Usecase.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Usecase.AddTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
