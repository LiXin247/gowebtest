package user

import (
	"github.com/gin-gonic/gin"
	"gowebtest/dbconnect"
	"gowebtest/model/user"
	"net/http"
)

func UserRegisterInit(context *gin.Context) {

}
func UserRegister(context *gin.Context) {
	name := context.PostForm("name")
	password := context.PostForm("password")
	ConfirmPassword := context.PostForm("ConfirmPassword")
	if password != ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Passwords do not match",
		})
		return
	}
	db := dbconnect.Database{}
	db.DatabaseConnect()
	var lastuser model.UserInfo
	db.DB.Last(&lastuser)
	UID := lastuser.UserId + 1
	var alreadyexists User
	result_1 := db.DB.Where("user_name = ?", name).First(&alreadyexists.Data)
	if result_1.Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The user already exists",
		})
		return
	}
	newUser := model.UserInfo{UID, name, password, 0, 0}
	result_2 := db.DB.Create(&newUser)
	if result_2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not create user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"user":    newUser,
	})
}
