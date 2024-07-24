package datacopy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func CardCopy(cardid int, cardimageurl string) (cardimageurlcopy string, err error) {
	// 获取图像的文件名
	imageFileName := strconv.Itoa(cardid) + ".jpg"
	// 获取当前工作目录
	projectPath, err := filepath.Abs(".")
	// 定义目标路径
	destPath := filepath.Join(projectPath, "cardimage", imageFileName) //绝对路径
	relativelyPath := filepath.Join("cardimage", imageFileName)        //相对路径
	if err != nil {
		fmt.Printf("Could not determine project path")
		return relativelyPath, err
	}
	// 打开源文件
	srcFile, err := os.Open(cardimageurl)
	if err != nil {
		fmt.Printf("Could not open source image\n")
		return relativelyPath, err
	}
	defer srcFile.Close()
	// 创建目标文件
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Could not create destination image\n")
		return relativelyPath, err
	}
	defer destFile.Close()
	// 复制文件内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {

		fmt.Printf("Could not datacopy image\n")
		return relativelyPath, err
	}
	// 返回成功响应
	fmt.Printf("Image uploaded successfully\n")
	return relativelyPath, err
}
