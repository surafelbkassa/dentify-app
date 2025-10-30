package main

import (
	"dentify/domain"
	"dentify/infrastructure"
	"dentify/repository"
	"dentify/usecase"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	connStr := "postgres://dentify_user:new_password@localhost:5432/dentify?sslmode=disable"
	fmt.Printf("Attempting to connect with: %s\n", connStr)

	db, err := infrastructure.ConnectDB(connStr)
	if err != nil {
		fmt.Printf("❌ Database connection FAILED: %v\n", err)
		panic(err)
	}
	fmt.Println("✅ Database connected successfully!")

	userRepo := &repository.PostgresUserRepo{DB: db}
	userUC := &usecase.UserUsecase{UserRepo: userRepo}

	r.POST("/signup", func(c *gin.Context) {
		var req domain.User
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Printf("❌ JSON bind error: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("📝 Attempting signup for user: %+v\n", req)

		if err := userUC.Signup(&req); err != nil {
			fmt.Printf("❌ Signup error: %s\n", err)
			msg := err.Error()
			switch {
			case strings.Contains(msg, "email") || strings.Contains(msg, "password") || strings.Contains(msg, "invalid") || strings.Contains(msg, "nil"):
				c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			case strings.Contains(msg, "already exists") || strings.Contains(msg, "duplicate"):
				c.JSON(http.StatusConflict, gin.H{"error": msg})
			default:
				fmt.Printf("🔥 Internal server error details: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		fmt.Printf("✅ Signup successful for user: %s\n", req.Email)
		req.Password = ""
		c.JSON(http.StatusCreated, gin.H{"message": "signup successful", "user": req})
	})
	r.POST("/login", func(c *gin.Context) {
		var loginReq struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userUC.Login(loginReq.Email, loginReq.Password)
		if err != nil {
			msg := err.Error()
			switch {
			case strings.Contains(msg, "email") || strings.Contains(msg, "password") || strings.Contains(msg, "invalid") || strings.Contains(msg, "nil"):
				c.JSON(http.StatusBadRequest, gin.H{"error": msg})
			case strings.Contains(msg, "not found"):
				c.JSON(http.StatusNotFound, gin.H{"error": msg})
			default:
				fmt.Printf("🔥 Internal server error details: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "login successful", "user": user})
	})

	r.Run(":8080")
}
