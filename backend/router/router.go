package router

import (
	"To-Do-Go/backend/controller"
	"To-Do-Go/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/login", controller.Login)
	r.GET("/api/todos", middleware.AuthorizeJWT(), controller.GetTodos)
	r.POST("/api/todos", middleware.AuthorizeJWT(), controller.AddTodo)

	return r
}
