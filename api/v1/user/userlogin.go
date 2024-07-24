package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowebtest/data/user"
	"net/http"
)

func UserLoginInit(context *gin.Context) {}
func UserLogin(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("password")
	err := user.UserLoginData(name, password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "Invalid credentials"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful", "user": name})
}
