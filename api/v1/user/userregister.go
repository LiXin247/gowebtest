package user

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data/user"
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
	err := user.UserRegisterData(name, password)
	switch err {
	case "The user already exists":
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The user already exists",
		})
		return
	case "Could not create user":
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not create user",
		})
		return
	case "nil":
		context.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User registered successfully",
			"user":    name,
		})

	}
}
