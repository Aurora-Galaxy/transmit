package control

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// FilesController 文件上传
func FilesController(c *gin.Context) {
	file, err := c.FormFile("raw") //raw前端参数名，从前端获取文件
	if err != nil {
		log.Fatal(err)
	}
	exe, err := os.Executable() //获取当前文件执行路径
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe) //获取当前文件的执行目录
	if err != nil {
		log.Fatal(err)
	}
	filename := uuid.New().String() //随机生成一个文件名
	uploads := filepath.Join(dir, "uploads")
	err = os.MkdirAll(uploads, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fullpath := path.Join("uploads", filename+filepath.Ext(file.Filename)) //文件存在本地的一个路径
	fileErr := c.SaveUploadedFile(file, filepath.Join(dir, fullpath))
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	c.JSON(http.StatusOK, gin.H{"url": "/" + fullpath})
}
