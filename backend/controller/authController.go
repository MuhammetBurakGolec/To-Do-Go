package controller

import (
	"To-Do-Go/backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Kullanıcı doğrulaması ve JWT token oluşturma işlemleri
	// Bu örnekte basit bir kontrol yapılıyor.
	// Gerçek bir uygulamada veritabanı sorgulamaları ve şifre doğrulaması gerekmektedir.

	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if loginInfo.Username == "admin" && loginInfo.Password == "password" {
		token, err := middleware.CreateToken(loginInfo.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating JWT token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
