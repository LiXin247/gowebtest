package card

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data/card"
	"net/http"
	"strconv"
)

func GachaCard(context *gin.Context) {
	userid, _ := strconv.Atoi(context.PostForm("userid"))
	frequency := context.PostForm("frequency")
	err := card.GachaCardData(frequency, userid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Gacha record failed",
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":    "successful",
			"frequency": "frequency",
		})
	}
}
