package user

import (
	"github.com/gin-gonic/gin"
	"gowebtest/model/datamodel"
)

type User struct {
	Data datamodel.UserInfo
}

func UserHome(context *gin.Context) {

}
