package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type admin struct {
	data gorm.Model
}

func AdminHome(context *gin.Context) {

}
