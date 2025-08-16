package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"todo_backend/internal/domain"
	"todo_backend/internal/infrastructure"
	"todo_backend/internal/infrastructure/mysql"
	"todo_backend/internal/interface/handler"
	"todo_backend/internal/usecase"
)

func main() {
	// .envを読み込む
	if err := godotenv.Load(".env"); err != nil {
		log.Println("[INFO] .env not found; using system environment variables")
	}

	// DB初期化（今回はSQLite）
	db, err := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbPath, _ := filepath.Abs("./todo.db")
	log.Println("USING_SQLITE:", dbPath)

	// マイグレーション
	if err := db.AutoMigrate(&domain.User{}, &domain.Todo{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// Repository
	userRepo := mysql.NewUserMySQL(db)
	todoRepo := mysql.NewTodoMysql(db)

	// Usecase
	authUC := usecase.NewAuthUsecase(userRepo)
	todoUC := usecase.NewTodoUsecase(todoRepo)

	// Handler
	authH := handler.NewAuthHandler(authUC)

	// ルータ生成
	router := infrastructure.NewRouter(authH, todoUC)

	// CORS追加
	router.Use(cors.Default())

	// JWT_SECRETチェック（開発中の注意喚起）
	if os.Getenv("JWT_SECRET") == "" {
		log.Println("[WARN] JWT_SECRET is not set. Set a strong secret in production.")
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
