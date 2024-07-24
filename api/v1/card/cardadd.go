package card

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gowebtest/data"
	"gowebtest/model/datamodel"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func CardAdd(context *gin.Context) {
	name := context.PostForm("name")
	cardlevel, err := strconv.Atoi(context.PostForm("cardlevel"))
	cardimageurl := context.PostForm("cardimageurl")
	description := context.PostForm("description")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid cardlevel",
		})
		return
	}
	db := data.Database{}
	db.DatabaseConnect()
	var alreadyexists Card
	result1 := db.DB.Where("card_name = ?", name).First(&alreadyexists.Data)
	if result1.Error != gorm.ErrRecordNotFound {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "The card already exists",
		})
		return
	}
	var lastcard datamodel.CardInfo
	db.DB.Last(&lastcard)
	CardID := lastcard.CardId + 1
	cardimageurlcopy, err_2 := CardCopy(CardID, cardimageurl, context)
	if err_2 != nil {
		return
	}
	newCard := datamodel.CardInfo{CardID, name, cardlevel, cardimageurlcopy, description}
	result2 := db.DB.Create(&newCard)
	if result2.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not create card",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "card add successfully",
		"user":    newCard,
	})

}
func CardCopy(cardid int, cardimageurl string, context *gin.Context) (cardimageurlcopy string, err error) {
	// 获取图像的文件名
	imageFileName := strconv.Itoa(cardid) + ".jpg"

	// 获取当前工作目录
	projectPath, err := filepath.Abs(".")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not determine project path",
		})
		return
	}

	// 定义目标路径
	destPath := filepath.Join(projectPath, "cardimage", imageFileName) //绝对路径
	relativelyPath := filepath.Join("cardimage", imageFileName)        //相对路径
	// 打开源文件
	srcFile, err := os.Open(cardimageurl)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not open source image",
		})
		return
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(destPath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not create destination image",
		})
		return
	}
	defer destFile.Close()

	// 复制文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Could not copy image",
		})
		return
	}

	// 返回成功响应
	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Image uploaded successfully",
		"path":    destPath,
	})
	return relativelyPath, err
}
