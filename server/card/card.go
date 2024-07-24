package card

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Card struct {
	Data gorm.Model
}

func Cards(context *gin.Context) {

}
