package handler

import (
	"net/http"
	"strconv"
	"todo_backend/internal/domain"
	"todo_backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// TodoHandlerは、HTTPリクエストとTodoユースケースをつなぐ役割を持つハンドラです。
// Ginのルーティングを通じて、Todoエンティティに関する操作を提供します。
type TodoHandler struct {
	Usecase *usecase.TodoUsecase
}

// NewTodoHandlerは、TodoHandlerを生成し、Ginのルーターにエンドポイントを登録します。
// r: Ginのエンジン
// uc: Todoユースケース
func NewTodoHandler(r *gin.Engine, uc *usecase.TodoUsecase) {
	h := &TodoHandler{Usecase: uc}
	r.GET("/todos", h.GetTodos)
	r.POST("/todos", h.CreateTodo)
	r.PUT("/todos/:id", h.UpdateTodo)
	r.DELETE("/todos/:id", h.DeleteTodo)
}

// GetTodosは、全てのTodoを取得してJSON形式で返します。
// HTTP:GET/todos
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.Usecase.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodoは、新しいTodoを作成します。
// リクエストボディはJSON形式で、Todo構造体にバインドされます。
// HTTP:POST/todos
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

// UpdateTodoは、既存のTodoを更新します。
// リクエストボディはJSON形式で、Todo構造体にバインドされます。
// HTTP:PUT/todos/:id
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Usecase.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteTodo は、指定されたIDのTodoを削除します。
// URLパラメータ:idを整数に変換して処理します。
// HTTP: DELETE /todos/:id
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.Usecase.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
