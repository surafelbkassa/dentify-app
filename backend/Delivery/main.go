package main

import (
	"dentify/domain"
	"dentify/infrastructure"
	"dentify/repository"
	"dentify/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := infrastructure.ConnectDB("postgres://dentify_user:securepassword@localhost:5432/dentify?sslmode=disable")
	if err != nil {
		panic(err)
	}
	userRepo := &repository.PostgresUserRepo{DB: db}
	userUC := &usecase.UserUsecase{UserRepo: userRepo}
	r.POST("/signup", func(c *gin.Context) {
		var req domain.User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := userUC.Signup(&req); err != nil {
			// map errors to HTTP statuses
			msg := err.Error()
			switch {
			case strings.Contains(msg, "email") || strings.Contains(msg, "password") || strings.Contains(msg, "invalid") || strings.Contains(msg, "nil"):
				c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			case strings.Contains(msg, "already exists") || strings.Contains(msg, "duplicate"):
				c.JSON(http.StatusConflict, gin.H{"error": msg})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		req.Password = ""
		c.JSON(http.StatusCreated, gin.H{"message": "signup successful", "user": req})
	})

	r.Run(":8080")
}
