package user

import (
	"github.com/gin-gonic/gin"
	"gowebtest/dbconnect"
	model "gowebtest/model/user"
	"net/http"
)

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
	newUser := model.UserInfo{UID, name, password, 0, 0}
	result := db.DB.Create(&(newUser))
	if result.Error != nil {
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
