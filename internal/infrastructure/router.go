package infrastructure

import (
	jwtmw "todo_backend/internal/infrastructure/jwt"
	"todo_backend/internal/interface/handler"
	"todo_backend/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(authHandler *handler.AuthHandler, todoUC *usecase.TodoUsecase) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// 認証不要
	r.POST("/signup", authHandler.Signup)
	r.POST("/login", authHandler.Login)

	// 認証必須のルート
	auth := r.Group("/")
	auth.Use(jwtmw.AuthRequired())
	{
		handler.NewTodoHandler(auth, todoUC)
	}

	return r
}
