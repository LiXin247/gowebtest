package card

import (
	"github.com/gin-gonic/gin"
	"gowebtest/data/card"
	"net/http"
	"strconv"
)

func CardAdd(context *gin.Context) {
	name := context.PostForm("name")
	cardlevel, err_1 := strconv.Atoi(context.PostForm("cardlevel"))
	cardimageurl := context.PostForm("cardimageurl")
	description := context.PostForm("description")
	if err_1 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid cardlevel",
		})
		return
	}
	err := card.CardAddData(name, cardlevel, cardimageurl, description)
	switch err {
	case "The card already exists":
		{
			context.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "The card already exists",
			})
			return
		}
	case "Card image copy failed":
		{
			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Card image copy failed",
			})
			return
		}
	case "Card image added failed":
		{
			context.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Card image added failed",
			})
			return
		}
	case "nil":
		{
			context.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "card add successfully",
				"name":    name,
			})
		}

	}

}
