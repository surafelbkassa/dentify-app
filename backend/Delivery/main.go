package delivery

import (
	"dentify/domain"
	"dentify/infrastructure"
	"dentify/repository"
	"dentify/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, _ := infrastructure.ConnectDB("postgres://user:12390ssD@localhost:5432/dentify")
	userRepo := &repository.PostgresUserRepo{DB: db}
	userUC := &usecase.UserUsecase{UserRepo: userRepo}
	r.POST("/signup", func(c *gin.Context) {
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := userUC.Signup(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "signup successful"})
	})

	r.Run(":8080")
}
