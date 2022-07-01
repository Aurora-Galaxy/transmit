package control

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// TextController 传递文本
func TextController(c *gin.Context){
	var json struct{
		Raw string `json:"raw"`
	}
	if err := c.ShouldBindJSON(&json) ; err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"err" : err})
	}else{
		exe , err := os.Executable()  //获取当前文件执行路径
		if err != nil {
			log.Fatal(err)
		}
		dir := filepath.Dir(exe) //获取当前文件的执行目录
		filename := uuid.New().String()  //随机生成一个文件名
		uploads := filepath.Join(dir,"uploads")
		err = os.MkdirAll(uploads , os.ModePerm) //创建uploads目录， os.ModePerm == 777 文件权限
		if err != nil{
			log.Fatal(err)
		}
		fullpath := path.Join("uploads",filename+".txt")
		err = ioutil.WriteFile(filepath.Join(dir,fullpath) , []byte(json.Raw),0644)
		if err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK , gin.H{"url":"/" + fullpath})  //返回文件的绝对路径
	}
}
