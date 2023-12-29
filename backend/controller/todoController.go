package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Basit bir Todo modeli (gerçek uygulamalarda veritabanı entegrasyonu gereklidir)
type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos = []Todo{
	{ID: 1, Task: "Alışveriş yap"},
	{ID: 2, Task: "Fatura öde"},
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func AddTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	todos = append(todos, newTodo)
	c.JSON(http.StatusOK, newTodo)
}
