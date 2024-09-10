package user

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data/user"
	"net/http"
)

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required"`
}

func UserRegisterInit(context *gin.Context) {

}
func UserRegister(context *gin.Context) {
	var req RegisterRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Name == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "name is null",
		})
		println("name is null")
		return
	}
	if req.Password != req.ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Passwords do not match",
		})
		return
	}
	err := user.UserRegisterData(req.Name, req.Password)
	switch err {
	case "The user already exists":
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The user already exists",
		})
		println("The user already exists")
		return
	case "Could not create user":
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not create user",
		})
		println("Could not create user")
		return
	case "nil":
		context.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User registered successfully",
			"user":    req.Name,
		})

	}
}
