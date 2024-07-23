package user

import (
	"github.com/gin-gonic/gin"
	model "gowebtest/model/user"
)

type user struct {
	data model.UserInfo
}

func UserHome(context *gin.Context) {

}
