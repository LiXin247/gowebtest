package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowebtest/dbconnect"
	"net/http"
)

func UserLoginInit(context *gin.Context) {

}
func UserLogin(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("password")
	db := dbconnect.Database{}
	db.DatabaseConnect()
	var u user
	result := db.DB.Where("user_name = ? AND user_password = ?", name, password).First(&u.data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized", "message": "Invalid credentials"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Internal server error"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "success", "message": "Login successful", "user": u.data})
}
