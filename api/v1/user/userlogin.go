package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowebtest/data/user"
	"net/http"
)

type LoginRuquest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func UserLoginInit(context *gin.Context) {}
func UserLogin(context *gin.Context) {
	var lrqe LoginRuquest
	if err := context.ShouldBindJSON(&lrqe); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := user.UserLoginData(lrqe.Name, lrqe.Password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "Invalid credentials"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful", "user": lrqe.Name})
}
