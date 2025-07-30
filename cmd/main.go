package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo_backend/internal/domain"
	"todo_backend/internal/infrastructure/mysql"
	"todo_backend/internal/interface/handler"
	"todo_backend/internal/usecase"
)

func main() {
	// DB初期化（今回はSQLite）
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// マイグレーション
	db.AutoMigrate(&domain.Todo{})

	// DI
	repo := mysql.NewTodoMysql(db)
	uc := usecase.NewTodoUsecase(repo)

	// Ginルーター設定
	r := gin.Default()

	// CORS追加
	r.Use(cors.Default())

	handler.NewTodoHandler(r, uc)

	r.Run(":8080")
}
